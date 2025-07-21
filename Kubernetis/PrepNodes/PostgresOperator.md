```bash
  git clone https://github.com/CrunchyData/postgres-operator-examples.git
  cd postgres-operator-examples
```
```bash
  kubectl apply -k kustomize/install/namespace
  kubectl apply --server-side -k kustomize/install/default
  kubectl -n postgres-operator get pods --selector=postgres-operator.crunchydata.com/control-plane=postgres-operator --field-selector=status.phase=Running
```
```bash
  kubectl apply -k kustomize/postgres
```
To get your Secrets 
```bash
  kubectl get secret hippo-pguser-hippo -n postgres-operator -o jsonpath="{.data}"
```
---> Then with .data.user  | base64 -d
    gives decoded Data
