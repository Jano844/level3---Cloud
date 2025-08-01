## Portforward

- Do all on Openstack Server in su - stack User

--Once to Activate Portforwarding
```bash
  sudo sysctl -w net.ipv4.ip_forward=1
  sudo nano /etc/sysctl.conf
  add "net.ipv4.ip_forward = 1" at buttom if not existintg
  sudo sysctl -p
```

--For Every Port you Want to Forward
```bash
  sudo iptables -t nat -A PREROUTING -i eth0 -p tcp --dport <Outside Port> -j DNAT --to-destination <Master/Worker IP>:<NodePort>
  sudo iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
  sudo iptables -A FORWARD -p tcp -d <Master/Worker IP> --dport <NodePort> -j ACCEPT
```

- OutsidePort = is Port that you want to use in Browser or Curl from Anywhere --> Can be chosen randomly like 2000

- NodePort = is the NodePort you choose , ```bash kubectl get service``` with the correct Namespace. Portrange 30000-32767

- Master/Worker IP = is floating Ip (in intenal net for example: 172.24.4.124) of your master Node
  if not working use ips of workers to check --> Kluster Comunication is not correct