variable "auth_url" {}
variable "tenant_name" {}
variable "user_name" {}
variable "password" {}
variable "region" {}

variable "k8s_node_count" {
  description = "Number of Kubernetes nodes (including control plane)"
  default     = 3
}

variable "k8s_flavor" {
  description = "OpenStack flavor for Kubernetes nodes"
  default     = "m1.medium"
}

variable "k8s_image" {
  description = "OpenStack image for Kubernetes nodes"
  default     = "ubuntu-22.04"
}

variable "k8s_network_name" {
  description = "OpenStack network name for Kubernetes nodes"
  default     = "private"
}

variable "k8s_external_network_name" {
  description = "OpenStack external network name for floating IPs"
  default     = "public"
}

variable "k8s_keypair" {
  description = "OpenStack keypair name for SSH access"
  default     = "default-key"
}

variable "k8s_public_key_path" {
  description = "Path to the SSH public key file to use for the keypair"
  default     = "~/.ssh/id_rsa.pub"
}