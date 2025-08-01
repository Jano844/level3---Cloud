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
  kubectl get secrets -n postgres-operator
  kubectl get secret <pguser> -n postgres-operator -o jsonpath="{.data}"
```
Get .data.password .data.user .data.host .data.dbname
all with | base64 -d && echo ""
Example:
```bash
  kubectl get secret my-postgres-cluster-pguser-myuser -n postgres-operator -o jsonpath="{.data.password}" | base64 -d && echo ""
```

```bash
  kubectl exec -it -n <pod name (4/4)> -- /bin/bash
  psql -h <host> -U <user> -d <database>
```
--> Now connected to Database



Show CRD
Edit CRD

```bash
  kubectl get crd postgresclusters.postgres-operator.crunchydata.com -o yaml
  kubectl edit crd postgresclusters.postgres-operator.crunchydata.com
```

CRD --> Defines how the CR muss look like
CR in this case in "/postgres-operator-examples/kustomize/postgres/postgres.yaml"
--> Defines how the Pod looks like (Change num of Instances or Username Password)

