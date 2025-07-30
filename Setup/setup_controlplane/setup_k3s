

-- Simple Setup with 1 Contolplane
Master:
curl -sfL https://get.k3s.io | sh -

get Token:
sudo cat /var/lib/rancher/k3s/server/node-token


Workers:
Change <Token> With the real Token
curl -sfL https://get.k3s.io | K3S_URL=https://172.24.4.101:6443 K3S_TOKEN=<Token> sh -


curl -sfL https://get.k3s.io | \
K3S_URL=https://172.24.4.140:6443 \
K3S_TOKEN="TOKEN" \
sh -

TEST: K3S_URL=https://192.168.151.115:6443 \

Check if it worked
sudo systemctl status k3s-agent
sudo journalctl -u k3s-agent -f

Master:
sudo kubectl get nodes

-- multiple Control Panels
Open Ports 
10250
6444
6443
2380
2379
(No Load Balancer)
curl -sfL https://get.k3s.io | sh -s - server --cluster-init --tls-san <Floating_ip_First_Node>
sudo cat /var/lib/rancher/k3s/server/node-token

Node 2 3 ....
curl -sfL https://get.k3s.io | K3S_TOKEN=<Token> sh -s - server --server https://<Floating_ip_First_Node>:6443 --tls-san <Floating_ip_First_Node>

sudo kubectl get nodes ## Check