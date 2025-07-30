terraform {
  required_providers {
    openstack = {
      source  = "terraform-provider-openstack/openstack"
      version = ">= 3.3.0"
    }
  }
}

provider "openstack" {
  # These variables should be set via environment variables or a .tfvars file
  auth_url    = var.auth_url
  tenant_name = var.tenant_name
  user_name   = var.user_name
  password    = var.password
  region      = var.region
}