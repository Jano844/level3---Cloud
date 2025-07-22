
// main.go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "os"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "myapp/createDB"
)

func main() {
    config, err := rest.InClusterConfig()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to get in-cluster config: %v\n", err)
        os.Exit(1)
    }

    var clientset *kubernetes.Clientset
    clientset, err = kubernetes.NewForConfig(config)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create clientset: %v\n", err)
        os.Exit(1)
    }


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "jano844/level3-test:v1.0.4\n")
	})

    http.HandleFunc("/pods", func(w http.ResponseWriter, r *http.Request) {
		handlePods(w, r, clientset)
	})

    http.HandleFunc("/createDB", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            createDB.CreateNewDatabase(w, r, clientset)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
	})

    http.ListenAndServe(":8080", nil)
}

// /pods Handler
func handlePods(w http.ResponseWriter, r *http.Request, clientset *kubernetes.Clientset) {
    pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        http.Error(w, "Failed to list pods: "+err.Error(), http.StatusInternalServerError)
        return
    }

    type PodInfo struct {
        Name   string `json:"name"`
        Status string `json:"status"`
    }

    var result []PodInfo
    for _, pod := range pods.Items {
        result = append(result, PodInfo{
            Name:   pod.Name,
            Status: string(pod.Status.Phase),
        })
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

