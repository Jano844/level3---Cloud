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
  name: postgres-cluster-{{ .Username }}
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
    - name: instance-{{ .Username }}
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
  name: postgres-cluster-{{ .Username }}-external
  namespace: postgres-operator
  labels:
    postgres-operator.crunchydata.com/cluster: postgres-cluster-{{ .Username }}
    app: postgres-external
spec:
  type: NodePort
  selector:
    postgres-operator.crunchydata.com/cluster: postgres-cluster-{{ .Username }}
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
	Username string `json:"username"`
	DbName   string `json:"dbname"`
}

// Template data structure
type ClusterTemplateData struct {
	Username  string
	Databases []string
	NodePort  int32
}

func CreateNewDatabase(w http.ResponseWriter, r *http.Request, clientset *kubernetes.Clientset) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Check Json
	var req DatabaseRequest
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

	// Check if PostgreSQL cluster already exists
	clusterName := fmt.Sprintf("postgres-cluster-%s", req.Username)
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
	templateData.Username = req.Username

	// Calculate NodePort (same for new and existing clusters)
	templateData.NodePort = int32(30000 + (len(clusterName) % 2767))

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
										"username":  req.Username,
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
		"username":         req.Username,
		"cluster":          clusterName,
		"databases":        templateData.Databases,
		"action":           action,
		"external_port":    templateData.NodePort,
		"external_service": fmt.Sprintf("%s-external", clusterName),
		"connection_info":  "External access enabled via NodePort",
	}

	json.NewEncoder(w).Encode(response)
}
