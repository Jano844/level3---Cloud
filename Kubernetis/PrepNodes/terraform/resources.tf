resource "openstack_compute_keypair_v2" "k8s_key" {
  name       = var.k8s_keypair
  public_key = file(var.k8s_public_key_path)
}

resource "openstack_networking_secgroup_v2" "k8s_secgroup" {
  name        = "k8s-secgroup"
  description = "Security group for Kubernetes nodes"
}

resource "openstack_networking_secgroup_rule_v2" "k8s_ssh" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = openstack_networking_secgroup_v2.k8s_secgroup.id
}

resource "openstack_networking_secgroup_rule_v2" "k8s_api" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 6443
  port_range_max    = 6443
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = openstack_networking_secgroup_v2.k8s_secgroup.id
}

resource "openstack_compute_instance_v2" "k8s_nodes" {
  count       = var.k8s_node_count
  name        = "k8s-node-${count.index + 1}"
  image_name  = var.k8s_image
  flavor_name = var.k8s_flavor
  key_pair    = openstack_compute_keypair_v2.k8s_key.name

  network {
    uuid = openstack_networking_network_v2.k8s_network.id
    port = openstack_networking_port_v2.k8s_ports[count.index].id
  }

  security_groups = [openstack_networking_secgroup_v2.k8s_secgroup.name]
}

# Create a port for each node
resource "openstack_networking_port_v2" "k8s_ports" {
  count             = var.k8s_node_count
  name              = "k8s-port-${count.index + 1}"
  network_id        = openstack_networking_network_v2.k8s_network.id
  security_group_ids = [openstack_networking_secgroup_v2.k8s_secgroup.id]

  fixed_ip {
    subnet_id = openstack_networking_subnet_v2.k8s_subnet.id
  }
}

# Assign floating IPs to each node
resource "openstack_networking_floatingip_v2" "k8s_fips" {
  count = var.k8s_node_count
  pool  = var.k8s_external_network_name
}

resource "openstack_networking_floatingip_associate_v2" "k8s_fip_assoc" {
  count       = var.k8s_node_count
  floating_ip = openstack_networking_floatingip_v2.k8s_fips[count.index].address
  port_id     = openstack_networking_port_v2.k8s_ports[count.index].id
}

output "k8s_node_ips" {
  value = [for node in openstack_compute_instance_v2.k8s_nodes : node.access_ip_v4]
}

output "k8s_node_floating_ips" {
  value = [for fip in openstack_networking_floatingip_v2.k8s_fips : fip.address]
}