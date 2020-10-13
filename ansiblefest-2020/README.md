# Ansiblefest 2020 - Centralized RBAC for Networks with Ansible Tower

This repository contains examples of automating using Ansible playbooks across network devices from multiple vendors including:

* Cisco
    * NXOS
    * IOS
* Arista
* VyOS
* Cumulus

The advantage of using Ansible to automate networking devices is that it provides a centralized toolset that can interact with different vendor devices simultaneously. Tower further extends the value of using Ansible for Network Automation by providing a built-in, encrypted credential store, SCM integration and full RBAC capabilities. 

# Ansible Tower Configuration as Code

[https://github.com/redhat-cop/tower_configuration](https://github.com/redhat-cop/tower_configuration)

The GNS3 Network Simulator is used to emulate the devices: 

[GNS3](https://www.gns3.com/software/download)

## Playbooks

| Playbook Name  | Description  |
|---|---|
| device_gather_facts.yaml  | Gather facts from multiple network devices  |
| network_report.yaml  | Gather facts from multiple network devices and produce a templated Jinja 2 report |
| cisco_nxos_declarative.yaml  | Configure Cisco NXOS Multi-Layer Switch from defined variable values  |
| ping.yaml  | Perform a network PING using vendor modules to verify Ansible can connect to network devices with valid credentials |