package createDB

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
)

// DeleteRequest structure for database deletion
type DeleteRequest struct {
	Username string `json:"username"`
	DbName   string `json:"dbname"`
}

func DeleteDatabase(w http.ResponseWriter, r *http.Request, clientset *kubernetes.Clientset) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Parse JSON request
	var req DeleteRequest
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

	// Check if PostgreSQL cluster exists
	clusterName := fmt.Sprintf("postgres-cluster-%s", req.Username)
	namespace := "postgres-operator"
	restClient := clientset.RESTClient()

	// Get existing cluster
	existingResult := restClient.
		Get().
		AbsPath("/apis/postgres-operator.crunchydata.com/v1beta1").
		Namespace(namespace).
		Resource("postgresclusters").
		Name(clusterName).
		Do(context.TODO())

	if existingResult.Error() != nil {
		response := map[string]interface{}{
			"error":    "Cluster not found",
			"cluster":  clusterName,
			"username": req.Username,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get current cluster configuration
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
	databaseFound := false

	if spec, ok := existingCluster["spec"].(map[string]interface{}); ok {
		if users, ok := spec["users"].([]interface{}); ok && len(users) > 0 {
			if user, ok := users[0].(map[string]interface{}); ok {
				if databases, ok := user["databases"].([]interface{}); ok {
					for _, db := range databases {
						if dbName, ok := db.(string); ok {
							if dbName == req.DbName {
								databaseFound = true
								// Don't add this database to the new list (effectively deleting it)
							} else {
								currentDatabases = append(currentDatabases, dbName)
							}
						}
					}
				}
			}
		}
	}

	// Check if database was found
	if !databaseFound {
		response := map[string]interface{}{
			"error":    "Database not found in cluster",
			"dbname":   req.DbName,
			"username": req.Username,
			"cluster":  clusterName,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if this was the last database
	if len(currentDatabases) == 0 {
		// Delete the entire cluster
		fmt.Printf("Deleting cluster '%s' (last database '%s' removed)\n", clusterName, req.DbName)

		result := restClient.
			Delete().
			AbsPath("/apis/postgres-operator.crunchydata.com/v1beta1").
			Namespace(namespace).
			Resource("postgresclusters").
			Name(clusterName).
			Do(context.TODO())

		if result.Error() != nil {
			http.Error(w, fmt.Sprintf("Failed to delete cluster: %v", result.Error()), http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"message":   "Successfully deleted database and cluster (last database)",
			"dbname":    req.DbName,
			"username":  req.Username,
			"cluster":   clusterName,
			"action":    "deleted_cluster",
			"databases": []string{},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Update cluster with remaining databases
	fmt.Printf("Removing database '%s' from cluster '%s'\n", req.DbName, clusterName)

	// Create template data with remaining databases
	templateData := ClusterTemplateData{
		Username:  req.Username,
		Databases: currentDatabases,
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

	fmt.Println("Generated YAML for update:")
	fmt.Println(yamlBuffer.String())

	var yamlObj map[string]interface{}
	if err := yaml.Unmarshal(yamlBuffer.Bytes(), &yamlObj); err != nil {
		fmt.Printf("Error parsing YAML: %v\n", err)
		http.Error(w, "Failed to parse generated YAML", http.StatusInternalServerError)
		return
	}

	// Preserve metadata (including resourceVersion) and update only the spec
	if yamlSpec, ok := yamlObj["spec"]; ok {
		existingCluster["spec"] = yamlSpec
	}

	jsonBytes, err := json.Marshal(existingCluster)
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
		return
	}

	// Update existing cluster with preserved metadata
	result := restClient.
		Put().
		AbsPath("/apis/postgres-operator.crunchydata.com/v1beta1").
		Namespace(namespace).
		Resource("postgresclusters").
		Name(clusterName).
		Body(jsonBytes).
		Do(context.TODO())

	if result.Error() != nil {
		http.Error(w, fmt.Sprintf("Failed to update database cluster: %v", result.Error()), http.StatusInternalServerError)
		return
	}

	// Return success response
	response := map[string]interface{}{
		"message":   "Successfully removed database from cluster",
		"dbname":    req.DbName,
		"username":  req.Username,
		"cluster":   clusterName,
		"action":    "database_removed",
		"databases": currentDatabases,
	}
	json.NewEncoder(w).Encode(response)
}
