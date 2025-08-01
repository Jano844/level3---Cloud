
-- Helm needed
```bash
    helm version
```

-- If helm not installed
```bash
    curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```


-- Add the Prometheus Community Helm Chart
```bash
    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
    helm repo update
    kubectl create namespace monitoring
```

-- If installation with K3s
```bash
    mkdir -p ~/.kube
    sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
    sudo chown $(id -u):$(id -g) ~/.kube/config
```

-- Install Prometheus and Grafana
```bash
    helm install prometheus-stack prometheus-community/kube-prometheus-stack --namespace monitoring
```

-- Get Grafana Admin Password
```bash
    kubectl --namespace monitoring get secrets prometheus-stack-grafana -o jsonpath="{.data.admin-password}" | base64 -d ; echo
```

-- Update Grafana Password
- Create file.yaml
```yaml
    grafana:
    adminPassword: "password"
    service:
        type: NodePort
        nodePort: 30001
```

```bash
    helm upgrade --install prometheus-stack prometheus-community/kube-prometheus-stack --namespace monitoring --create-namespace -f filename.yaml
```