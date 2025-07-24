# Go Backend API Documentation

This Go backend provides APIs for managing PostgreSQL databases in Kubernetes using the Postgres Operator.

## API Base URL
```
http://localhost:30800
```
*(Replace with your actual cluster IP and NodePort)*

## Available Endpoints

### 1. Health Check
**Endpoint:** `GET /`

**Description:** Check if the API is running

**Example:**
```bash
curl -X GET http://localhost:30800/
```

**Response:**
```
jano844/level3-test:v1.1.4
```

---

### 2. List All Pods
**Endpoint:** `GET /pods`

**Description:** List all pods across all namespaces

**Example:**
```bash
curl -X GET http://localhost:30800/pods
```

**Response:**
```json
[
  {
    "name": "postgres-cluster-john-instance1-abcd",
    "namespace": "postgres-operator",
    "status": "Running"
  }
]
```

---

### 3. Create Database
**Endpoint:** `POST /createDB`

**Description:** Create a new PostgreSQL database with automatic external NodePort access

**Example:**
```bash
curl -X POST http://localhost:30800/createDB \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john",
    "dbname": "mydatabase"
  }'
```

**Response:**
```json
{
  "message": "Successfully created database cluster",
  "dbname": "mydatabase",
  "username": "john",
  "cluster": "postgres-cluster-john",
  "databases": ["mydatabase"],
  "action": "created",
  "external_port": 30021,
  "external_service": "postgres-cluster-john-ha",
  "connection_info": "External access enabled via NodePort"
}
```

**Adding additional database to existing cluster:**
```bash
curl -X POST http://localhost:30800/createDB \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john",
    "dbname": "anotherdatabase"
  }'
```

---

### 4. Get Database Access Information
**Endpoint:** `POST /getDBAccess`

**Description:** Get connection information for a specific database

**Example:**
```bash
curl -X POST http://localhost:30800/getDBAccess \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john",
    "dbname": "mydatabase"
  }'
```

**Response:**
```json
{
  "status": "success",
  "message": "Database access information retrieved successfully",
  "database_name": "mydatabase",
  "username": "john",
  "cluster_name": "postgres-cluster-john",
  "internal_host": "postgres-cluster-john-primary.postgres-operator.svc.cluster.local",
  "internal_port": 5432,
  "external_host": "YOUR_CLUSTER_IP",
  "external_port": 30021,
  "password": "*** Use: kubectl get secret postgres-cluster-john-pguser-john -n postgres-operator -o jsonpath='{.data.password}' | base64 -d ***",
  "connection_string": "postgresql://john:[PASSWORD]@postgres-cluster-john-primary.postgres-operator.svc.cluster.local:5432/mydatabase?sslmode=require",
  "external_connection_string": "postgresql://john:[PASSWORD]@YOUR_CLUSTER_IP:30021/mydatabase?sslmode=require"
}
```

---

### 5. Delete Database
**Endpoint:** `DELETE /deleteDB`

**Description:** Delete a PostgreSQL cluster

**Example:**
```bash
curl -X DELETE http://localhost:30800/deleteDB \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john",
    "dbname": "mydatabase"
  }'
```

**Response:**
```json
{
  "message": "Successfully deleted database cluster",
  "cluster": "postgres-cluster-john"
}
```

---

## Complete Workflow Example

### 1. Create a new database with external access:
```bash
curl -X POST http://localhost:30800/createDB \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "dbname": "testdb"
  }'
```

### 2. Get connection information:
```bash
curl -X POST http://localhost:30800/getDBAccess \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "dbname": "testdb"
  }'
```

### 3. Get the database password:
```bash
kubectl get secret postgres-cluster-testuser-pguser-testuser -n postgres-operator \
  -o jsonpath='{.data.password}' | base64 -d
```

### 4. Connect to the database from outside the cluster:
```bash
# Get cluster node IP
kubectl get nodes -o wide

# Connect (replace CLUSTER_IP and PASSWORD)
psql "postgresql://testuser:PASSWORD@CLUSTER_IP:30021/testdb?sslmode=require"
```

### 5. Verify external service was created:
```bash
kubectl get services -n postgres-operator --field-selector spec.type=NodePort
```

### 6. Clean up when done:
```bash
curl -X DELETE http://localhost:30800/deleteDB \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "dbname": "testdb"
  }'
```

---

## Error Responses

### Bad Request (400):
```json
{
  "error": "Username and dbname are required"
}
```

### Conflict (409) - Database already exists:
```json
{
  "message": "Database already exists in cluster",
  "dbname": "mydatabase",
  "username": "john",
  "cluster": "postgres-cluster-john",
  "databases": ["mydatabase"]
}
```

### Not Found (404):
```json
{
  "status": "error",
  "message": "PostgreSQL cluster 'postgres-cluster-john' not found"
}
```

---

## Features

✅ **Automatic NodePort Creation** - External access is configured automatically when creating databases  
✅ **Multi-Database Support** - Add multiple databases to the same cluster  
✅ **Connection Information** - Get both internal and external connection details  
✅ **CORS Enabled** - Can be called from web frontends  
✅ **Error Handling** - Proper HTTP status codes and error messages  

## Requirements

- Kubernetes cluster with Postgres Operator installed
- Go backend running in the cluster with proper RBAC permissions
- Access to the cluster from outside (for NodePort connections)

## Notes

- NodePorts are automatically assigned in the range 30000-32767
- External IP needs to be replaced with actual cluster node IP
- Passwords must be retrieved using kubectl commands
- Services are created in the `postgres-operator` namespace
