- Overview: 2 Vms (nodes)
--> Openstack Network

- Master Node --> API Server Controlls Workers
- Workers --> Run Application Pods (Docker Container)


          [ Kubernetes Cluster ]
                   |
     +-------------------------------+
     |         Virtual Network       |
     |         (OpenStack)           |
     +-------------------------------+
        |                     |
+---------------+    +----------------+
|  Master Node  |    |  Worker Node   |
| (K8s Control) |    | (Pods laufen)  |
+---------------+    +----------------+
                          |
                    +-------------+
                    |  Pod A      |
                    |  Container  |
                    +-------------+
                    |  Pod B      |
                    +-------------+

