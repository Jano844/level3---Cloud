
import openstack

# Verbindung herstellen
conn = openstack.connect(cloud='mycloud')

# 1. Netzwerk erstellen
network = conn.network.create_network(name='k8s-network')
subnet = conn.network.create_subnet(
    name='k8s-subnet',
    network_id=network.id,
    ip_version=4,
    cidr='192.168.100.0/24',
    gateway_ip='192.168.100.1',
    dns_nameservers=['8.8.8.8']
)

# 2. Router erstellen und Subnetz verbinden
ext_net = conn.network.find_network('public')  # oder dein externer Netzwerkname
router = conn.network.create_router(name='k8s-router', external_gateway_info={'network_id': ext_net.id})
conn.network.add_interface_to_router(router, subnet_id=subnet.id)

# 3. Keypair erstellen (nur falls noch nicht vorhanden)
keypair_name = 'k8s-key'
if not conn.compute.find_keypair(keypair_name):
    with open('~/.ssh/id_rsa.pub') as pubkey_file:
        conn.compute.create_keypair(name=keypair_name, public_key=pubkey_file.read())

# 4. Security Group mit Regeln
secgroup = conn.network.create_security_group(name='k8s-secgroup')
rules = [
    dict(protocol='tcp', port_range_min=22, port_range_max=22),
    dict(protocol='tcp', port_range_min=6443, port_range_max=6443),
    dict(protocol='tcp', port_range_min=30000, port_range_max=32767),
    dict(protocol='icmp')
]
for rule in rules:
    conn.network.create_security_group_rule(
        security_group_id=secgroup.id,
        direction='ingress',
        **rule
    )

# 5. VMs erstellen (Master + Worker)
image = conn.compute.find_image('ubuntu-22.04')
flavor = conn.compute.find_flavor('m1.medium')
net = conn.network.find_network('k8s-network')

for name in ['k8s-master', 'k8s-worker-1']:
    conn.compute.create_server(
        name=name,
        image_id=image.id,
        flavor_id=flavor.id,
        networks=[{"uuid": net.id}],
        key_name=keypair_name,
        security_groups=[{"name": secgroup.name}]
    )

