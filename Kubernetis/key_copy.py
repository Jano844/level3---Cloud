import os
import paramiko
from scp import SCPClient


def copy_key_to_stack():
    key_name = input("Enter the name of the key (without .pem extension): ").strip()
    local_key_path = f"{key_name}.pem"
    
    if not os.path.exists(local_key_path):
        print(f"Error: {local_key_path} not found in current directory")
        return

    server_ip = "91.99.176.18"
    print(f"Copying {local_key_path} to stack user on {server_ip}")
    
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    
    try:
        password = input("Enter Password: ").strip()
        ssh.connect(
            server_ip,
            username='stack',
            password=password,
            timeout=10
        )
        print("✅ Successfully connected as 'stack' user.")

        stdin, stdout, stderr = ssh.exec_command("mkdir -p ~/.ssh && chmod 700 ~/.ssh")
        exit_status = stdout.channel.recv_exit_status()
        if exit_status != 0:
            print(f"Failed to create Dir {stderr.read().decode()}")
            return

        with SCPClient(ssh.get_transport()) as scp:
            scp.put(local_key_path, f".ssh/{key_name}")

        ssh.exec_command(f"chmod 600 ~/.ssh/{key_name}")

        print("🎉 Successfully deployed key!")
        
    except Exception as e:
        print(f" Unexpected error: {str(e)}")
    finally:
        ssh.close()

if __name__ == "__main__":
    copy_key_to_stack()
