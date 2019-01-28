# Ansible Tower in Azure Cloud

The Azure cloud provides the infrastructure as a service for this Ansible Tower build. 

# Overview

This role handles infra and install/setup of a sole Ansible Tower in Azure cloud.
It uses a CentOS image. A license file must be provided post-install.
This infra-code handles:

- building the Azure infra
- preparing the VM for Tower
- runs the Tower installer
- inserts a license file if provided

It does not handle:

- preparing DNS/ FQDN for Tower .. specific to our use-case, otherwise it can be automated.


The structure of the module is as follows:
```
tower
├── tower.inv.j2
├── ansible.cfg
├── provision.yml
├── README.md
├── inventory
│   ├── hosts
│   └── host_vars
│       └── localhost
│           └── general.yml
├── roles
│   ├── azure-tower
│   │   ├── defaults
│   │   │   └── main.yml
│   │   ├── handlers
│   │   │   └── main.yml
│   │   ├── meta
│   │   │   └── main.yml
│   │   ├── README.md
│   │   ├── tasks
│   │   │   ├── configure.yml
│   │   │   └── main.yml
│   │   └── vars
│   │       └── main.yml
│   └── tower-install
│       ├── defaults
│       │   └── main.yml
│       ├── files
│       │   └── license
│       ├── handlers
│       │   └── main.yml
│       ├── meta
│       │   └── main.yml
│       ├── README.md
│       ├── tasks
│       │   └── main.yml
│       ├── templates
│       │   └── tower.inv.j2
│       ├── tests
│       │   ├── inventory
│       │   └── test.yml
│       └── vars
│           └── main.yml
└── vars
    └── extra_vars.yml

```

# Initial Setup

Prepare the following files with your own information

* extra_vars.yml

Login to Azure and perform the following steps.

```
az login    # browser opens and login can happen.. (not entirely necessary, but convenient for debugging)
# export environment variables to allow Azure authentication
export {ARM_CLIENT_ID,TF_VAR_ARM_CLIENT_ID}=<insert yours here>
export {ARM_CLIENT_SECRET,TF_VAR_ARM_CLIENT_SECRET}=<insert yours here>
export {ARM_SUBSCRIPTION_ID,TF_VAR_ARM_SUBSCRIPTION_ID}=<insert yours here>
export {ARM_TENANT_ID,TF_VAR_ARM_TENANT_ID}=<insert yours here>

```
You will need to prepare the DNS part separately as it's not handled within this project.

Run the playbook to make the role happen.

```
ansible-playbook -i localhost provision.yml
```

Feel free to submit PRs or Issues for improvements.
