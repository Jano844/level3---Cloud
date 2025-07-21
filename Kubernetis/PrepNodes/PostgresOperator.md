bash ```

git clone https://github.com/CrunchyData/postgres-operator-examples.git

cd postgres-operator-examples

kubectl apply -k kustomize/install/namespace
kubectl apply --server-side -k kustomize/install/default
kubectl -n postgres-operator get pods --selector=postgres-operator.crunchydata.com/control-plane=postgres-operator --field-selector=status.phase=Running

kubectl apply -k kustomize/postgres



  
```
