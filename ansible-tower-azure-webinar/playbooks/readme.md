# Ansible Demo Playbooks for Azure

## Background

These playbooks are designed to spin up isolated "lab" environments in Azure, handy for dev playground or student environments.

If we have a look at the inventory file, each group will set up a `resource group` in Azure that has it's own `vnet`, `subnet`, and `security group`. VM's in that group will be created in the `resouce group`. A `public IP` will be created for each VM to allow ssh access. 

```
[lab1]
lab1-vm01
lab1-vm02

[lab2]
lab2-vm01
lab2-vm02

[lab1:vars]
lab_name=lab1

[lab2:vars]
lab_name=lab2
```
Having a look at the Ansible Role structure the `defaults/main.yaml` defines the variables we will be using. Update the variables for naming, subnet, vm size, etc. as needed.

In the `files/` directory we have our public and private key for ssh to the Azure VMs. The public key will be pushed to the VM when Ansible creates it. The private key is encrypted in Ansible Vault. You'll want to delete these keys, generate another pair, `ssh-keygen`, and encrypt the private key `ansible-vault encypt roles/azure-lab/files/id_rsa`. Encryption is done so you don't have open private keys in your git repo. When you want to ssh to the VM you can decrypt the key and ssh. 

```
├── azure-lab.yaml
├── inventory
├── readme.md
└── roles
    └── azure-lab
        ├── defaults
        │   └── main.yaml
        ├── files
        │   ├── id_rsa
        │   └── id_rsa.pub
        └── tasks
            ├── create-infra.yaml
            ├── create-vm.yaml
            ├── destroy-lab.yaml
            ├── destroy-vm.yaml
            ├── main.yaml
            ├── shutdown-vm.yaml
            └── start-vm.yaml
```

The `/tasks` folder contains the playbooks to run, we toggle what playbook is running by setting the `activity` variable:
```
ansible-playbook -i inventory azure-lab.yaml -e "activity=create-infra"
ansible-playbook -i inventory azure-lab.yaml -e "activity=create-vm"
ansible-playbook -i inventory azure-lab.yaml -e "destroy-lab"
```

## Running the Playbooks
This will go over using the `ansible-playbook` command with these `.yaml` files. Once this is running from the command line you can set up the playbooks in Tower. 

To set up in Tower you'll need to create the project (point to repo with this code), the inventory (sourced from the project), azure credentials (using your SP account in Azure AD), and job template (pointing at the `azure-lab.yaml` in the project).

Make sure you have ansible installed on whatever OS you are using. You'll also want to have Azure cli installed and Azure SDK for python to run the Azure modules.


* Azure CLI install:

https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest

* SDK Install for Azure modules in playbooks

`sudo pip install ansible[azure]`

Authenticate to Azure before running playbooks, this can be done with `az login` and then set your  subscription in Azure with `az account set --subscription ca7XXXXX`

Finally run the playbooks!

```
ansible-playbook -i inventory azure-lab.yaml -e "activity=create-infra"
ansible-playbook -i inventory azure-lab.yaml -e "activity=create-vm"
```
When the the `create-infra` activity is done you'll get output about the public IP.
```
TASK [azure-lab : Print public IP for VMs] *************************************
ok: [lab2-vm01] => {
    "msg": "The public IP is 40.85.247.231."
}
```

You can now ssh in. First decrypt your private key.
```
ansible-vault decrypt roles/azure-lab/files/id_rsa
ssh -i roles/azure-lab/files/id_rsa azureadmin@40.85.247.231
```
## Helpful Links

https://github.com/Azure-Samples/ansible-playbooks

https://docs.microsoft.com/en-us/azure/ansible/

https://docs.ansible.com/ansible/latest/scenario_guides/guide_azure.html
