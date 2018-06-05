#!/bin/bash
terraform init
terraform apply -auto-approve
TF_STATE=terraform.tfstate ansible-playbook -i $(which terraform-inventory) site.yml