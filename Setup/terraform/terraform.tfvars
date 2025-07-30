auth_url    = "http://88.99.80.96/identity/v3"
tenant_name = "admin"
user_name   = "admin"
password    = "password"
region      = "RegionOne"

k8s_node_count   = 2
k8s_flavor       = "custom.v1"
k8s_image        = "Ubuntu 22.04 Jammy"
k8s_network_name = "private"
k8s_keypair      = "default-key"
k8s_public_key_path = "~/.ssh/id_rsa.pub"