
package createDB

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"k8s.io/client-go/kubernetes"
)

// Template für die PostgreSQL-Cluster-YAML
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
        - {{ .DbName }}
  instances:
    - name: instance-{{ .Username }}
      dataVolumeClaimSpec:
        accessModes:
        - "ReadWriteOnce"
        resources:
          requests:
            storage: 1Gi
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
`

// Struktur für die JSON-Daten
type DatabaseRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}


func CreateNewDatabase(w http.ResponseWriter, r *http.Request, clientset *kubernetes.Clientset) {

	// Check Json
	var req DatabaseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	tmpl, err := template.New("postgresCluster").Parse(postgresClusterTemplate)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	var yamlBuffer bytes.Buffer
	err = tmpl.Execute(&yamlBuffer, req)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}

	fmt.Println("Generated YAML:")
	fmt.Println(yamlBuffer.String())

  // Somewhere here is the Error cat parse Yaml to Json but YAML Correct
  
	var jsonData map[string]interface{}
	err = json.Unmarshal(yamlBuffer.Bytes(), &jsonData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to convert YAML to JSON: %v", err), http.StatusInternalServerError)
		return
	}

	namespace := "postgres-operator"
	restClient := clientset.RESTClient()

	// Ressource erstellen
	result := restClient.
		Post().
		AbsPath("/apis/postgres-operator.crunchydata.com/v1beta1").
		Namespace(namespace).
		Resource("postgresclusters").
		Body(jsonData).
		Do(context.TODO())

	// Error
	if result.Error() != nil {
		http.Error(w, fmt.Sprintf("Failed to create database: %v", result.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully created database: %s", req.DbName)
}
