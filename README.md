# level3---Cloud


# 🚀 OpenStack Kernkomponenten – Übersicht

Diese Datei bietet eine kompakte und verständliche Einführung in die vier zentralen OpenStack-Komponenten: **Nova**, **Neutron**, **Glance** und **Keystone**. Jede dieser Komponenten erfüllt eine spezifische Rolle im Cloud-Ökosystem.

---

## 🔧 Nova – Compute Service

### 🧭 Zweck  
Nova ist die **Compute-Komponente** von OpenStack. Sie ist verantwortlich für das **Starten und Verwalten virtueller Maschinen (VMs)**.

### ⚙️ Wichtige Funktionen
- Starten, Stoppen, Pausieren und Löschen von VMs.
- Verwaltung von Ressourcen wie CPU, RAM und Festplatten.
- Unterstützung verschiedener Hypervisoren (z. B. KVM, VMware, Xen).
- Orchestrierung der Lebenszyklen von Instanzen.
- Skalierung und Verfügbarkeitsmanagement von Compute-Ressourcen.

### 💡 Beispiel  
Ein Benutzer möchte eine neue VM starten. Nova nimmt die Anfrage entgegen, kommuniziert mit dem Hypervisor und weist der VM Speicher und CPU zu.

---

## 🌐 Neutron – Networking Service

### 🧭 Zweck  
Neutron ist die **Netzwerk-Komponente** von OpenStack. Sie stellt Netzwerkinfrastruktur und -dienste bereit.

### ⚙️ Wichtige Funktionen
- Verwaltung von Netzwerken, Subnetzen und Ports.
- IP-Zuweisung via DHCP.
- Unterstützung für VLANs, VXLANs und andere Overlay-Netzwerke.
- Integration mit SDN-Controllern (z. B. Open vSwitch, OVN).
- Nutzung von Floating IPs (öffentliche IPs für VMs).
- Sicherheitsgruppen zur Port-basierten Zugriffskontrolle (Firewall).

### 💡 Beispiel  
Ein Benutzer möchte zwei VMs in einem privaten Netzwerk verbinden und ihnen eine öffentliche IP zuweisen. Neutron regelt die interne Verbindung und das Routing nach außen.

---

## 🖼️ Glance – Image Service

### 🧭 Zweck  
Glance ist der **Image-Dienst** von OpenStack. Er verwaltet **VM-Images**, die als Vorlage für neue Instanzen dienen.

### ⚙️ Wichtige Funktionen
- Hochladen, Suchen, Löschen und Verwalten von Images.
- Unterstützung gängiger Image-Formate wie QCOW2, RAW, VMDK.
- Verwaltung von Image-Metadaten.
- Nutzung verschiedener Backends: Swift, Ceph, lokale Speicherung u.a.

### 💡 Beispiel  
Ein Administrator lädt ein Ubuntu-Image in Glance hoch. Nova kann dieses Image dann verwenden, um neue Ubuntu-Instanzen zu starten.

---

## 🔐 Keystone – Identity Service

### 🧭 Zweck  
Keystone ist die zentrale **Authentifizierungs- und Autorisierungskomponente** von OpenStack.

### ⚙️ Wichtige Funktionen
- Authentifizierung von Benutzern über verschiedene Methoden (Passwort, Token, LDAP, Federation).
- Rollenbasierte Zugriffskontrolle (RBAC).
- Bereitstellung eines Service-Katalogs mit allen verfügbaren OpenStack-Diensten.
- Token-Ausgabe zur sicheren Kommunikation mit APIs.

### 💡 Beispiel  
Ein Benutzer meldet sich über das OpenStack-Dashboard (Horizon) an. Keystone validiert die Zugangsdaten, gibt ein Token zurück und erlaubt damit den Zugriff auf Ressourcen.

---

## 🧩 Zusammenspiel der Komponenten

> **Beispielablauf beim Start einer VM:**

1. 🔐 **Keystone** authentifiziert den Benutzer und gibt ein API-Token aus.  
2. 🖼️ **Glance** liefert das gewünschte Image.  
3. 🌐 **Neutron** stellt Netzwerkressourcen (Netzwerk, IPs) bereit.  
4. 🔧 **Nova** orchestriert die Erstellung der VM mithilfe der bereitgestellten Ressourcen.

---

📘 Weitere Informationen findest du in der offiziellen [OpenStack-Dokumentation](https://docs.openstack.org/).

