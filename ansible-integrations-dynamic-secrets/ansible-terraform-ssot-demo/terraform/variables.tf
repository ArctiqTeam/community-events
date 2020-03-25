variable "credentials" {
  type = string
  description = "File pathname to GCP credentials."
}

variable "project" {
  default = "ansible-terraform-ssot-demo"
  description = "GCP project ID."
}

variable "region" {
  default = "northamerica-northeast1"
  description = "GCP project region."
}

variable "compute_address_name" {
  default = "ipv4-address"
  description = "Name of the external IP address."
}

variable "vm_name" {
  default = "vm1"
  description = "VM instance name."
}

variable "vm_machine_type" {
  default = "n1-standard-1"
  description = "VM machine type."
}

variable "vm_zone" {
  default = "northamerica-northeast1-c"
  description = "GCP zone for VM to be created in."
}

variable "vm_tags" {
  default = ["dev"]
  description = "Tags to be attached to the VM instance."
}

variable "gce_ssh_user" {
  default = "markieta"
  description = "User account to create and copy SSH key to."
}

variable "gce_ssh_pub_key" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDB29rHJa7vx32XJZ3BwyIvZ4b5hsqOYfkRIPeeu/oJE/+tymkpPfpEAV+7dB8OLfhsb+JQAju1jb/CYSijeZo5wu4FlmWmo2VIX9WRGHjEl4LpX6oatPqqXVmtdPN03EvQ0hutTIn1j9+YSA+0+Ke8MI3peuiAlgfbW+bTQNZ9K2fhH0Sg5oujrJXUHjsDB3S+js4CjCu6kViJT+5geX5S74FRo9Mu8eHfDX2dUxa9+2TaL9XSNWqY0OrJIZX/AVKtC21oodZ30BQD3y80HT1k7ARrN3ySM95yGsa7KcniqEDSpX2+5h6YcT5JIYuS/tdeNKs4pGFMXUkw9krGygF9VH2gETuWSuZSMHMu9ZJ5Wq7saYrHeuFQ6gdi3+zBsaHOKOC55Lu1ndiwq210oit9r913sy9ISBGJCO5p74tQfZcvaVwUOot9a90Bh4LZo0U+0sLacS4m5P2UKhoy3tzg2tB6tX7YF0jlKF35fQeMa3goWJU4cGS+6WH6YK41CVs= markieta@x1c"
  description = "SSH public key to copy to new VM."
}

variable "vm_boot_disk_image" {
  default = "centos-cloud/centos-7"
  description = "OS installation image."
}

variable "vm_network_interface_network" {
  default = "default"
  description = "Network name to be attached to VM instance's interface."
}

variable "ansible_inventory_dest" {
  default = "../ansible/inventories/templated/hosts"
  description = "Destination for Terraform to output templated inventory file."
}

variable "ansible_command" {
  default = "ANSIBLE_CONFIG='../ansible/ansible.cfg' ansible-playbook -i "
  description = "Ansible-playbook command and option to read inventory."
}

variable "ansible_playbook" {
  default = "../ansible/prompt.yml"
  description = "Path to Ansible playbook."
}


variable "tower_curl" {
  default = "curl -f -k -H 'Content-Type: application/json' -XPOST --user "
  description = "Curl command and options to request job launch on Tower."
}

variable "tower_username" {
  default = "null"
  description = "Tower username to launch job."
}

variable "tower_password" {
  default = "null"
  description = "Tower password to launch job."
}

variable "tower_launch_url" {
  default = "https://tower/api/v2/workflow_job_templates/14/launch/"
  description = "URL to Tower API to launch specific job."
}

variable "initiator" {
  default = "terraform"
  description = "Application that is initiating this run."
}
