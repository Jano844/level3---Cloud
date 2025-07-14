import os
import paramiko
from scp import SCPClient
import getpass

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
        # Connect as 'stack' with password
        password = input("Enter Password: ").strip()
        ssh.connect(
            server_ip,
            username='stack',
            password=password,
            timeout=10  # Prevents hanging
        )
        print("✅ Successfully connected as 'stack' user.")

        # Ensure .ssh directory exists
        stdin, stdout, stderr = ssh.exec_command("mkdir -p ~/.ssh && chmod 700 ~/.ssh")
        exit_status = stdout.channel.recv_exit_status()  # Wait for command to finish
        if exit_status != 0:
            print(f"⚠️ Failed to create .ssh directory: {stderr.read().decode()}")
            return

        # Copy the key using SCP (relative to stack's home directory)
        with SCPClient(ssh.get_transport()) as scp:
            scp.put(local_key_path, f".ssh/{key_name}")
        print(f"🔑 Key copied to /home/stack/.ssh/{key_name}")

        # Set strict permissions (required for SSH)
        ssh.exec_command(f"chmod 600 ~/.ssh/{key_name}")
        print("🔒 Permissions set to 600.")

        print("🎉 Successfully deployed key!")
        
    except Exception as e:
        print(f" Unexpected error: {str(e)}")
    finally:
        ssh.close()

if __name__ == "__main__":
    copy_key_to_stack()
