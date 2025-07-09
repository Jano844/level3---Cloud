

Diese Anleitung zeigt dir, wie du mit DevStack eine virtuelle Maschine (VM) in OpenStack erstellst und dich per SSH verbindest.



OpenStack Umgebung konfigurieren**  
Führe das OpenRC-Skript aus, um Umgebungsvariablen wie Benutzername, Projekt und Passwort zu setzen:  

bash
source ~/devstack/openrc admin admin


Das sorgt dafür, dass die OpenStack-CLI weiß, mit wem sie kommunizieren soll.


Ubuntu Cloud-Image herunterladen
Lade das offizielle Ubuntu Jammy Cloud-Image herunter:

bash
Kopieren
Bearbeiten
wget https://cloud-images.ubuntu.com/jammy/20250523/jammy-server-cloudimg-amd64.img
Dieses Image dient als Betriebssystem für deine VM.


Image in OpenStack registrieren
Importiere das heruntergeladene Image in OpenStack:

bash
Kopieren
Bearbeiten
openstack image create "Ubuntu 22.04 Jammy" \
  --file jammy-server-cloudimg-amd64.img \
  --disk-format qcow2 \
  --container-format bare
qcow2 ist das bevorzugte Format für Cloud-Images, bare bedeutet, das Image hat keinen Container.


SSH-Schlüssel erzeugen
Erzeuge ein Schlüsselpaar für die SSH-Verbindung:

bash
Kopieren
Bearbeiten
ssh-keygen -t rsa -f ~/.ssh/mykey
Mit ls ~/.ssh/mykey* kannst du die Dateien prüfen.


Keypair zu OpenStack hinzufügen
Importiere den öffentlichen Schlüssel in OpenStack:

bash
Kopieren
Bearbeiten
openstack keypair create --public-key ~/.ssh/mykey.pub mykey
So kannst du dich später per SSH anmelden.


Flavors prüfen oder erstellen
Schau dir vorhandene Flavors an:

bash
Kopieren
Bearbeiten
openstack flavor list
Falls kein passender dabei ist, erstelle einen (z.B. 1GB RAM, 1 vCPU, 10 GB Festplatte):

bash
Kopieren
Bearbeiten
openstack flavor create flavor_01 --ram 1024 --disk 10 --vcpus 1

Security Group prüfen
Liste alle Security Groups auf:

bash
Kopieren
Bearbeiten
openstack security group list

Security Group für SSH anlegen
Erstelle eine neue Security Group, die SSH erlaubt:

bash
Kopieren
Bearbeiten
openstack security group create ssh-access --description "Allow SSH from anywhere"
Erlaube den Zugriff auf Port 22 von überall:

bash
Kopieren
Bearbeiten
openstack security group rule create ssh-access --protocol tcp --dst-port 22 --remote-ip 0.0.0.0/0


VM erstellen
Erstelle die VM mit deinem Flavor, Image, Netzwerk, Keypair und Security Group:

bash
Kopieren
Bearbeiten
openstack server create ubuntu-test-vm \
  --flavor ds1G \
  --image "Ubuntu 22.04 Jammy" \
  --nic net-id=$(openstack network show private -f value -c id) \
  --key-name mykey \
  --security-group ssh-access
Ersetze ds1G ggf. durch einen Flavor, der in deinem System existiert.



Floating IP erstellen und zuweisen
Erstelle eine Floating IP aus dem öffentlichen Pool:

bash
Kopieren
Bearbeiten
openstack floating ip create public
Weise sie deiner VM zu (ersetze <floating-ip> durch die tatsächliche IP):

bash
Kopieren
Bearbeiten
openstack server add floating ip ubuntu-test-vm <floating-ip>
Prüfe, ob die IP zugeordnet wurde:

bash
Kopieren
Bearbeiten
openstack server show ubuntu-test-vm -f value -c addresses


Mit der VM verbinden
Verbinde dich per SSH:

bash
Kopieren
Bearbeiten
ssh -i ~/.ssh/mykey ubuntu@<floating-ip>
Wenn die Verbindung nicht klappt, prüfe das Konsolen-Log der VM:

bash
Kopieren
Bearbeiten
openstack console log show ubuntu-test-vm | tail -n 50


Damit hast du erfolgreich eine VM in OpenStack mit DevStack aufgesetzt und kannst dich darauf einloggen.
