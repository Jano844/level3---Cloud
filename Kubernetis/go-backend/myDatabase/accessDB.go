package myDatabase

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"k8s.io/client-go/kubernetes"
)

// DatabaseAccessRequest structure for API requests
type DatabaseAccessRequest struct {
	Username string `json:"username"`
	DbName   string `json:"dbname"`
}

// DatabaseAccessResponse structure for API responses
type DatabaseAccessResponse struct {
	Status                   string `json:"status"`
	Message                  string `json:"message"`
	DatabaseName             string `json:"database_name"`
	Username                 string `json:"username"`
	ClusterName              string `json:"cluster_name"`
	InternalHost             string `json:"internal_host"`
	InternalPort             int32  `json:"internal_port"`
	ExternalHost             string `json:"external_host"`
	ExternalPort             int32  `json:"external_port"`
	Password                 string `json:"password"`
	ActualPassword           string `json:"actual_password"`
	PSQLCommand              string `json:"psql_command"`
	ConnectionString         string `json:"connection_string"`
	ExternalConnectionString string `json:"external_connection_string"`
}

// GetDatabaseAccess returns connection information for a database
func GetDatabaseAccess(w http.ResponseWriter, r *http.Request, clientset *kubernetes.Clientset) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request
	var req DatabaseAccessRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Username == "" || req.DbName == "" {
		http.Error(w, "Username and dbname are required", http.StatusBadRequest)
		return
	}

	clusterName := fmt.Sprintf("postgres-cluster-%s", req.Username)
	namespace := "postgres-operator"

	// Check if cluster exists
	restClient := clientset.RESTClient()
	existingResult := restClient.
		Get().
		AbsPath("/apis/postgres-operator.crunchydata.com/v1beta1").
		Namespace(namespace).
		Resource("postgresclusters").
		Name(clusterName).
		Do(context.TODO())

	if existingResult.Error() != nil {
		response := DatabaseAccessResponse{
			Status:  "error",
			Message: fmt.Sprintf("PostgreSQL cluster '%s' not found", clusterName),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if cluster has the requested database
	var existingCluster map[string]interface{}
	rawCluster, err := existingResult.Raw()
	if err != nil {
		http.Error(w, "Failed to read cluster data", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(rawCluster, &existingCluster)
	if err != nil {
		http.Error(w, "Failed to parse cluster data", http.StatusInternalServerError)
		return
	}

	// Verify database exists in cluster
	databaseExists := false
	if spec, ok := existingCluster["spec"].(map[string]interface{}); ok {
		if users, ok := spec["users"].([]interface{}); ok && len(users) > 0 {
			if user, ok := users[0].(map[string]interface{}); ok {
				if databases, ok := user["databases"].([]interface{}); ok {
					for _, db := range databases {
						if dbName, ok := db.(string); ok && dbName == req.DbName {
							databaseExists = true
							break
						}
					}
				}
			}
		}
	}

	if !databaseExists {
		response := DatabaseAccessResponse{
			Status:  "error",
			Message: fmt.Sprintf("Database '%s' not found in cluster '%s'", req.DbName, clusterName),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get internal service information
	internalHost := fmt.Sprintf("%s-primary.%s.svc.cluster.local", clusterName, namespace)

	// Get external service information using proper REST client
	externalServiceName := fmt.Sprintf("%s-ha", clusterName)

	var externalPort int32 = 0
	var externalHost string = "Not configured"

	// Get the service using REST client - proper endpoint
	serviceResult := restClient.
		Get().
		AbsPath("/api/v1").
		Namespace(namespace).
		Resource("services").
		Name(externalServiceName).
		Do(context.TODO())

	if serviceResult.Error() == nil {
		// External service exists, get the NodePort
		rawService, err := serviceResult.Raw()
		if err == nil {
			var service map[string]interface{}
			if json.Unmarshal(rawService, &service) == nil {
				if spec, ok := service["spec"].(map[string]interface{}); ok {
					// Check if it's a NodePort service
					if serviceType, ok := spec["type"].(string); ok && serviceType == "NodePort" {
						if ports, ok := spec["ports"].([]interface{}); ok {
							// Look for the PostgreSQL port (5432)
							for _, portInterface := range ports {
								if port, ok := portInterface.(map[string]interface{}); ok {
									if portNum, ok := port["port"].(float64); ok && int32(portNum) == 5432 {
										if nodePort, ok := port["nodePort"].(float64); ok {
											externalPort = int32(nodePort)
											externalHost = "YOUR_CLUSTER_IP" // User needs to replace this
											break
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// Get actual password from secret
	secretName := fmt.Sprintf("%s-pguser-%s", clusterName, req.Username)
	secretResult := restClient.
		Get().
		AbsPath("/api/v1").
		Namespace(namespace).
		Resource("secrets").
		Name(secretName).
		Do(context.TODO())

	var actualPassword string = "Could not retrieve password"
	if secretResult.Error() == nil {
		rawSecret, err := secretResult.Raw()
		if err == nil {
			var secret map[string]interface{}
			if json.Unmarshal(rawSecret, &secret) == nil {
				if data, ok := secret["data"].(map[string]interface{}); ok {
					if passwordB64, ok := data["password"].(string); ok {
						// Decode base64 password
						if passwordBytes, err := base64.StdEncoding.DecodeString(passwordB64); err == nil {
							actualPassword = string(passwordBytes)
						}
					}
				}
			}
		}
	}

	// Password placeholder (user needs kubectl to get actual password)
	password := "*** Use: kubectl get secret " + clusterName + "-pguser-" + req.Username + " -n " + namespace + " -o jsonpath='{.data.password}' | base64 -d ***"

	// Generate psql command
	psqlCommand := fmt.Sprintf("PGPASSWORD='%s' psql -h <IP> -p %d -U %s -d %s",
		actualPassword, externalPort, req.Username, req.DbName)

	// Build connection strings
	internalConnectionString := fmt.Sprintf("postgresql://%s:[PASSWORD]@%s:5432/%s?sslmode=require",
		req.Username, internalHost, req.DbName)

	var externalConnectionString string
	if externalPort > 0 {
		externalConnectionString = fmt.Sprintf("postgresql://%s:[PASSWORD]@%s:%d/%s?sslmode=require",
			req.Username, externalHost, externalPort, req.DbName)
	} else {
		externalConnectionString = "External access not configured"
	}

	// Return success response
	response := DatabaseAccessResponse{
		Status:                   "success",
		Message:                  "Database access information retrieved successfully",
		DatabaseName:             req.DbName,
		Username:                 req.Username,
		ClusterName:              clusterName,
		InternalHost:             internalHost,
		InternalPort:             5432,
		ExternalHost:             externalHost,
		ExternalPort:             externalPort,
		Password:                 password,
		ActualPassword:           actualPassword,
		PSQLCommand:              psqlCommand,
		ConnectionString:         internalConnectionString,
		ExternalConnectionString: externalConnectionString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
