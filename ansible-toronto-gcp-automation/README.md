## Ansible and the Google Cloud Platform
The purpose of this demo is to show that you can control the state of your Compute Instances inside your Google Cloud projects with Ansible. With this Ansible playbook, users are able to control which Compute Instances should exist and which ones shouldn't, all controlled with a single inventory file. 

### How the Playbook Works
The playbook, *site.yml* , runs on a group of hosts called `compute-instances` and its first task is to query all Compute Instances that currently exist in a certain project and then compiles the names into an Ansible list called `existing_instance_list`. Once we have our list of current Compute Instances, two tasks are run, one task creates an Ansible group that will contain the new hosts that need to be provisioned (the delta of the `compute-instances` hosts and Compute Instances that currently exist in GCP) and the other task also creates an Ansible group that will contain the hosts that need to be destroyed (compilation of Compute instances that currently exist in GCP and hosts that are not in `compute-instances`).

Once we lists are compiled, the playbook runs a play on our newely created `new_instances` group that will provision any new Compute Instances via the `gcp_compute_instance` module. The playbook then takes cares of destroying nodes by running a final play on our newely created `old_instances` group that will destroy any nodes we no longer want in our Google project.

### Requirements
* Ansible >= 2.7
* python >= 2.6
* requests >= 2.18.4
* google-auth >= 1.3.0

In order to utilize the Google Cloud Ansible modules, you must specify a service account that has the permissions to read and write Compute Instances. Once a service account is created, generate a JSON file that will contain all authentication necessary to iteract with the GCP API and put the JSON file in the same working directory as the inventory file.

### Example Usecase
For this usecase, lets assume I currently have 2 Compute Instances in my GCP project: 
* instance-1
* instance-2

I now want to add a third node to my project called `instance-3` with the following settings:
* Running Ubuntu 16.04
* 50 GB disk
* n1-standard-1 machine

Our next step is to define our above attributes in a variable file for our new host:
```
~]# cat << EOF host_vars/instance-3
---
source_image: projects/ubuntu-os-cloud/global/images/family/ubuntu-1604-lts
disk_size_gb: 50
network_interface_name: eth0
machine_type: n1-standard-1
EOF
```

As a reference, my directory structure for this project currently looks like this:
```
~]# tree
.
├── README.md
├── group_vars
│   └── all
├── host_vars
│   └── instance-3
├── inventory
├── service.json
└── site.yml

2 directories, 6 files
```
Once we have our variable file for the new node, we need to add it to our inventory file. The inventory file should look like this:
```
~]# cat inventory
[compute-instances]
instance-1
instance-2
instance-3

[new_instances]

[old_instances]
```

The final step is to run the playbook and this will create our new Compute Instance `instance-3`:
```
~]# ansible-playbook -i inventory site.yml
```

Now lets assume we want to destroy that third Compute Instance that was just created. This is as simple as removing the name from the inventory file and re-running the same playbook. Our inventory file now looks like this:
```
~]# cat inventory
[compute-instances]
instance-1
instance-2

[new_instances]

[old_instances]
```

And we run the playbook:
```
~]# ansible-playbook -i inventory site.yml
```

