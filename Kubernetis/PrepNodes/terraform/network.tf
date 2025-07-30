resource "openstack_networking_network_v2" "k8s_network" {
  name           = var.k8s_network_name
  admin_state_up = true
}

resource "openstack_networking_subnet_v2" "k8s_subnet" {
  name            = "k8s-subnet"
  network_id      = openstack_networking_network_v2.k8s_network.id
  cidr            = "192.168.100.0/24"
  ip_version      = 4
  dns_nameservers = ["1.1.1.1", "8.8.8.8"]
}


data "openstack_networking_network_v2" "external_network" {
  name = var.k8s_external_network_name
}

resource "openstack_networking_router_v2" "k8s_router" {
  name                = "k8s-router"
  external_network_id = data.openstack_networking_network_v2.external_network.id
}

resource "openstack_networking_router_interface_v2" "k8s_router_interface" {
  router_id = openstack_networking_router_v2.k8s_router.id
  subnet_id = openstack_networking_subnet_v2.k8s_subnet.id
}