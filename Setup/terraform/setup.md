-- Get Cloud Image
```bash
wget https://cloud-images.ubuntu.com/jammy/20250619/jammy-server-cloudimg-amd64-disk-kvm.img

openstack image create "Ubuntu 22.04 Jammy" \
   --file jammy-server-cloudimg-amd64-disk-kvm.img \
   --disk-format qcow2 \
   --container-format bare
```

-- Create Costum Flavor
```bash
openstack flavor create costum.v1 \
    --ram 16384 \
    --disk 60 \
    --vcpus 8 \
    --public
```

-- Run terraform
```bash
    terraform init
    terraform plan
    terraform apply
```

-- Delete Instaces
```bash
  terraform
```


-- Place SSH KEYS
```bash
    su - stack
    mkdir -p ~/.ssh
    touch ~/.ssh/id_rsa
    touch ~/.ssh/id_rsa.pub
```
-- Copie Your Key
    Private and public Key from ~/.ssh/"your Key"
    the ones you used in terraform.tfvars

```bash
    chmod 400  ~/.ssh/id_rsa
    chmod 400  ~/.ssh/id_rsa.pub
```

-- Done

