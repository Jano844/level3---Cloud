import openstack
import random

# CP Yaml to where teh SDk can Find it
# mkdir -p /root/.config/openstack
# cp /app/config/clouds.yaml /root/.config/openstack/clouds.yaml



def generate_cidr():
    third_octet = random.randint(100, 200)  # oder anderer Bereich
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
            print(f"Error ddeleting Network: {e}")

    print("Network and Subnet Were Cretaed")
    print(f"Network: {name_net}")
    print(f"Subnet: {name_sub}")


def security_group_and_ssh(conn):
    pass

def master(conn):
    pass

def worker(conn):
    pass


def main():
    conn = openstack.connect(cloud='mycloud')

    network_and_subnet(conn)
    # for project in conn.identity.projects():
    #     print(f"- Projekt: {project.name}")


if __name__ == "__main__":
    main()
