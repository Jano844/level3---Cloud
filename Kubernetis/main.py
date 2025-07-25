import openstack
import random
import os
from key_copy import copy_key_to_stack

# CP Yaml to where teh SDk can Find it
# mkdir -p /root/.config/openstack
# cp /app/config/clouds.yaml /root/.config/openstack/clouds.yaml


def generate_cidr():
    third_octet = random.randint(100, 200)
    return f"192.168.{third_octet}.0/24"

def network_and_subnet(conn):

    """
    This Function ascs you in the Commandline
    for a name of a network and then creates it

    Same happens for the subnet, it will laso create you a random CDIR
    and if it works prints it out 
    """

    # Network
    try:
        name_net : str = input("Enter a network name: ").strip()
        network = conn.network.create_network(name=name_net)
    except Exception as e:
        print("Error Creating Network")
        print(f"Details: {e}")
        return


    #subnet
    try:
        name_sub : str = input("Enter a subnet name: ").strip()
        cidr = generate_cidr()

        subnet = conn.network.create_subnet(
            name=name_sub,
            network_id=network.id,
            ip_version='4',
            cidr=cidr,
            gateway_ip=cidr.replace("0/24", "1"),
            dns_nameservers=['8.8.8.8'],
            enable_dhcp=True
        )
    except Exception as e:
        print("Error Creating Subnet")
        print(f"Details: {e}")
        try:
            conn.network.delete_network(network.id)
        except Exception as e:
            print(f"Error deleting Network: {e}")

    print("Network and Subnet Were Cretaed")
    print(f"Network: {name_net}")
    print(f"Subnet: {name_sub}")

    return name_net, name_sub




def security_group(conn):
    try:
        name = input("Enter a name for the Security Group: ").strip()

        # Create the security group
        secgroup = conn.network.create_security_group(name=name)

        # Add SSH (Port 22) rule
        conn.network.create_security_group_rule(
            security_group_id=secgroup.id,
            direction='ingress',
            protocol='tcp',
            port_range_min=22,
            port_range_max=22,
            remote_ip_prefix='0.0.0.0/0',
            ethertype='IPv4'
        )

        conn.network.create_security_group_rule(
            security_group_id=secgroup.id,
            direction='ingress',
            protocol='icmp',
            remote_ip_prefix='0.0.0.0/0',
            ethertype='IPv4'
        )

        return secgroup

    except Exception as e:
        print("Failed to create security group or rules")
        print(f"Details: {e}")


def ssh_pair(conn):
    try:
        name = input("Enter a name for the SSH keypair: ").strip()
        keypair = conn.compute.create_keypair(name=name)

        private_key = keypair.private_key
        if private_key:
            filename = f"{name}.pem"
            with open(filename, "w") as f:
                f.write(private_key)
            os.chmod(filename, 0o600)
            print(f"Private key saved to {filename}")
        else:
            print("No private key returned (key might already exist)")

        return keypair

    except Exception as e:
        print("Failed to create keypair")
        print(f"Details: {e}")





def assign_floating_ip_correctly(conn, vm_name):
    ext_net = next(net for net in conn.network.networks() if net.is_router_external)
    
    server = conn.compute.find_server(vm_name)
    if not server:
        print(f"[ERROR] Server '{vm_name}' not found.")
        return

    ports = list(conn.network.ports(device_id=server.id))
    if not ports:
        print(f"[ERROR] No ports found for server {vm_name}")
        return

    internal_port = ports[0]
    print(f"[INFO] Using port: {internal_port.id}, fixed IP: {internal_port.fixed_ips[0]['ip_address']}")

    fip = conn.network.create_ip(
        floating_network_id=ext_net.id,
        port_id=internal_port.id,
    )

    print(f"[SUCCESS] Floating IP {fip.floating_ip_address} assigned to {vm_name}")






def launch_instance(conn, key, security_group_name, net_name):

    name = input("Enter instance name: ").strip()

    image = conn.compute.find_image("Ubuntu 22.04 Jammy")
    flavor = conn.compute.find_flavor("ds1G")
    network = conn.network.find_network(net_name)

    if not image:
        print("Image not found!")
        return
    if not flavor:
        print("Flavor not found!")
        return


    server = conn.compute.create_server(
        name=name,
        image_id=image.id,
        flavor_id=flavor.id,
        networks=[{"uuid": network.id}],
        key_name=key.name,
        security_groups=[{"name": security_group_name}]
    )

    server = conn.compute.wait_for_server(server)
    print(f"✅ VM '{name}' is ready! IPs: {server.addresses}")
    return server



def main():
    servers = []
    conn = openstack.connect(cloud='mycloud')

    net_name, name_sub = network_and_subnet(conn)
    sec_group = security_group(conn)
    key = ssh_pair(conn)


    # router = conn.network.find_router('router1')
    router = conn.network.get_router("93b07f98-e60b-4748-9aa1-a30d3239af1f")
    subnet = conn.network.find_subnet(name_sub)
    conn.network.add_interface_to_router(router, subnet_id=subnet.id)

    for _ in range(3):
        server = launch_instance(conn, key, sec_group.name, net_name)
        floating_ip = assign_floating_ip_correctly(conn, server.name)
        servers.append(server)
        print(floating_ip)


    copy_key_to_stack()

   


if __name__ == "__main__":
    main()
