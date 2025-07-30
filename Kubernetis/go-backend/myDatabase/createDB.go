package myDatabase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Template für die PostgreSQL-Cluster-YAML mit NodePort Service
const postgresClusterTemplate = `
apiVersion: postgres-operator.crunchydata.com/v1beta1
kind: PostgresCluster
metadata:
  name: pgc-{{ .Username }}
  namespace: postgres-operator
  annotations:
    postgres-operator.crunchydata.com/autoCreateUserSchema: "true"
spec:
  postgresVersion: 17
  users:
    - name: {{ .Username }}
      databases:
        {{- range .Databases }}
        - {{ . }}
        {{- end }}
  instances:
    - name: inst-{{ .Username }}
      dataVolumeClaimSpec:
        accessModes:
        - "ReadWriteOnce"
        resources:
          requests:
            storage: 2Gi
  backups:
    pgbackrest:
      repos:
      - name: repo1
        volume:
          volumeClaimSpec:
            accessModes:
            - "ReadWriteOnce"
            resources:
              requests:
                storage: 1Gi
  service:
    type: NodePort
    nodePort: {{ .NodePort }}
---
apiVersion: v1
kind: Service
metadata:
  name: pgc-{{ .Username }}-external
  namespace: postgres-operator
  labels:
    postgres-operator.crunchydata.com/cluster: pgc-{{ .Username }}
    app: postgres-external
spec:
  type: NodePort
  selector:
    postgres-operator.crunchydata.com/cluster: pgc-{{ .Username }}
    postgres-operator.crunchydata.com/role: master
  ports:
  - name: postgres
    port: 5432
    targetPort: 5432
    nodePort: {{ .NodePort }}
    protocol: TCP
`

// Struktur für die JSON-Daten
type DatabaseRequest struct {
	DbName string `json:"dbname"`
}

// Template data structure
type ClusterTemplateData struct {
	Username  string
	Databases []string
	NodePort  int32
}

// findAvailableNodePort finds an available NodePort using smart collision detection
func findAvailableNodePort(restClient rest.Interface, namespace string, username string, isExistingCluster bool) (int32, error) {
	// If it's an existing cluster, try to get the current NodePort from the existing service
	if isExistingCluster {
		expectedServiceName := fmt.Sprintf("pgc-%s-ha", username)

		serviceResult := restClient.
			Get().
			AbsPath("/api/v1").
			Namespace(namespace).
			Resource("services").
			Name(expectedServiceName).
			Do(context.TODO())

		if serviceResult.Error() == nil {
			rawService, err := serviceResult.Raw()
			if err == nil {
				var service map[string]interface{}
				if json.Unmarshal(rawService, &service) == nil {
					if spec, ok := service["spec"].(map[string]interface{}); ok {
						if ports, ok := spec["ports"].([]interface{}); ok && len(ports) > 0 {
							if port, ok := ports[0].(map[string]interface{}); ok {
								if nodePort, ok := port["nodePort"].(float64); ok {
									fmt.Printf("DEBUG: Reusing existing NodePort %d for user %s\n", int32(nodePort), username)
									return int32(nodePort), nil
								}
							}
						}
					}
				}
			}
		}
	}

	// Get all existing NodePort services to find available port
	servicesResult := restClient.
		Get().
		AbsPath("/api/v1").
		Namespace(namespace).
		Resource("services").
		Do(context.TODO())

	if servicesResult.Error() != nil {
		return 0, fmt.Errorf("failed to list services: %v", servicesResult.Error())
	}

	rawServices, err := servicesResult.Raw()
	if err != nil {
		return 0, fmt.Errorf("failed to read services: %v", err)
	}

	var servicesList map[string]interface{}
	if err := json.Unmarshal(rawServices, &servicesList); err != nil {
		return 0, fmt.Errorf("failed to parse services list: %v", err)
	}

	// Collect all used NodePorts
	usedPorts := make(map[int32]bool)
	if items, ok := servicesList["items"].([]interface{}); ok {
		for _, item := range items {
			if service, ok := item.(map[string]interface{}); ok {
				if spec, ok := service["spec"].(map[string]interface{}); ok {
					if serviceType, ok := spec["type"].(string); ok && serviceType == "NodePort" {
						if ports, ok := spec["ports"].([]interface{}); ok {
							for _, portInterface := range ports {
								if port, ok := portInterface.(map[string]interface{}); ok {
									if nodePort, ok := port["nodePort"].(float64); ok {
										usedPorts[int32(nodePort)] = true
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("DEBUG: Found %d used NodePorts\n", len(usedPorts))

	// Start with a simple hash-based port for better distribution, then search for free port
	usernameHash := 0
	for _, char := range username {
		usernameHash = usernameHash*31 + int(char)
	}
	if usernameHash < 0 {
		usernameHash = -usernameHash
	}
	startPort := int32(30000 + (usernameHash % 2768))

	// Try to find available port starting from hash-based port
	maxAttempts := 2768 // Maximum possible ports in range
	currentPort := startPort

	for attempt := 0; attempt < maxAttempts; attempt++ {
		if !usedPorts[currentPort] {
			fmt.Printf("DEBUG: Found available NodePort %d for user %s (attempt %d, started from %d)\n",
				currentPort, username, attempt+1, startPort)
			return currentPort, nil
		}

		// Try next port
		currentPort++
		if currentPort > 32767 {
			currentPort = 30000 // Wrap around to beginning of NodePort range
		}
	}

	return 0, fmt.Errorf("no available NodePort found after checking all %d possible ports. NodePort range exhausted", maxAttempts)
}

func CreateNewDatabase(w http.ResponseWriter, r *http.Request, clientset *kubernetes.Clientset) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Extract username from JWT token context
	username, ok := r.Context().Value("user_id").(string)
	if !ok || username == "" {
		http.Error(w, "Unable to get user information from token", http.StatusUnauthorized)
		return
	}
	username, err := ToAlphaBaseStr(username)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	// Check Json
	var req DatabaseRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.DbName == "" {
		http.Error(w, "dbname is required", http.StatusBadRequest)
		return
	}

	// Check if PostgreSQL cluster already exists
	clusterName := fmt.Sprintf("pgc-%s", username)
	namespace := "postgres-operator"
	restClient := clientset.RESTClient()

	// Try to get existing cluster
	existingResult := restClient.
		Get().
		AbsPath("/apis/postgres-operator.crunchydata.com/v1beta1").
		Namespace(namespace).
		Resource("postgresclusters").
		Name(clusterName).
		Do(context.TODO())

	var templateData ClusterTemplateData
	templateData.Username = username

	// Find available NodePort by checking existing services
	availablePort, err := findAvailableNodePort(restClient, namespace, username, existingResult.Error() == nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find available NodePort: %v", err), http.StatusInternalServerError)
		return
	}

	templateData.NodePort = availablePort
	fmt.Printf("DEBUG: Assigned NodePort %d to user '%s'\n", availablePort, username)

	if existingResult.Error() == nil {
		// Cluster exists - get current databases and add new one
		var existingCluster map[string]interface{}
		rawCluster, err := existingResult.Raw()
		if err != nil {
			http.Error(w, "Failed to read existing cluster", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(rawCluster, &existingCluster)
		if err != nil {
			http.Error(w, "Failed to parse existing cluster", http.StatusInternalServerError)
			return
		}

		// Extract current databases
		currentDatabases := make([]string, 0)
		if spec, ok := existingCluster["spec"].(map[string]interface{}); ok {
			if users, ok := spec["users"].([]interface{}); ok && len(users) > 0 {
				if user, ok := users[0].(map[string]interface{}); ok {
					if databases, ok := user["databases"].([]interface{}); ok {
						for _, db := range databases {
							if dbName, ok := db.(string); ok {
								currentDatabases = append(currentDatabases, dbName)
								// Check if database already exists
								if dbName == req.DbName {
									response := map[string]interface{}{
										"message":   "Database already exists in cluster",
										"dbname":    req.DbName,
										"username":  username,
										"cluster":   clusterName,
										"databases": currentDatabases,
									}
									w.WriteHeader(http.StatusConflict)
									json.NewEncoder(w).Encode(response)
									return
								}
							}
						}
					}
				}
			}
		}

		// Add new database to existing list
		templateData.Databases = append(currentDatabases, req.DbName)

		// Update existing cluster
		fmt.Printf("Adding database '%s' to existing cluster '%s'\n", req.DbName, clusterName)
	} else {
		// Create new cluster with first database
		templateData.Databases = []string{req.DbName}
		fmt.Printf("Creating new cluster '%s' with database '%s'\n", clusterName, req.DbName)
	}

	// Generate YAML with updated database list
	tmpl, err := template.New("postgresCluster").Parse(postgresClusterTemplate)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	var yamlBuffer bytes.Buffer
	err = tmpl.Execute(&yamlBuffer, templateData)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}

	fmt.Println("Generated YAML:")
	fmt.Println(yamlBuffer.String())

	var yamlObj map[string]interface{}
	if err := yaml.Unmarshal(yamlBuffer.Bytes(), &yamlObj); err != nil {
		fmt.Printf("Fehler beim Parsen von YAML: %v\n", err)
		http.Error(w, "Failed to parse generated YAML", http.StatusInternalServerError)
		return
	}

	// Create or update the PostgreSQL cluster
	var result rest.Result
	if existingResult.Error() == nil {
		// For updates, we need to get the existing resource and merge our changes
		var existingCluster map[string]interface{}
		rawCluster, err := existingResult.Raw()
		if err != nil {
			http.Error(w, "Failed to read existing cluster for update", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(rawCluster, &existingCluster)
		if err != nil {
			http.Error(w, "Failed to parse existing cluster for update", http.StatusInternalServerError)
			return
		}

		// Preserve metadata (including resourceVersion) and update only the spec
		if yamlSpec, ok := yamlObj["spec"]; ok {
			existingCluster["spec"] = yamlSpec
		}

		jsonBytes, err := json.Marshal(existingCluster)
		if err != nil {
			fmt.Printf("Fehler beim Konvertieren zu JSON: %v\n", err)
			http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
			return
		}

		// Update existing cluster with preserved metadata
		result = restClient.
			Put().
			AbsPath("/apis/postgres-operator.crunchydata.com/v1beta1").
			Namespace(namespace).
			Resource("postgresclusters").
			Name(clusterName).
			Body(jsonBytes).
			Do(context.TODO())
	} else {
		// Create new cluster
		jsonBytes, err := json.Marshal(yamlObj)
		if err != nil {
			fmt.Printf("Fehler beim Konvertieren zu JSON: %v\n", err)
			http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
			return
		}

		result = restClient.
			Post().
			AbsPath("/apis/postgres-operator.crunchydata.com/v1beta1").
			Namespace(namespace).
			Resource("postgresclusters").
			Body(jsonBytes).
			Do(context.TODO())
	}

	// Check for errors
	if result.Error() != nil {
		http.Error(w, fmt.Sprintf("Failed to create/update database cluster: %v", result.Error()), http.StatusInternalServerError)
		return
	}

	// Determine action for response
	action := "created"
	if existingResult.Error() == nil {
		action = "updated"
	}

	// Return success response as JSON
	response := map[string]interface{}{
		"message":          fmt.Sprintf("Successfully %s database cluster", action),
		"dbname":           req.DbName,
		"username":         username,
		"cluster":          clusterName,
		"databases":        templateData.Databases,
		"action":           action,
		"external_port":    templateData.NodePort,
		"external_service": fmt.Sprintf("%s-external", clusterName),
		"connection_info":  "External access enabled via NodePort",
	}

	json.NewEncoder(w).Encode(response)
}
