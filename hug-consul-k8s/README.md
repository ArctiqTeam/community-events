# This is a demo on Consul and K8s integration - Hashicorp User Group Oct 2018

The demonstration during the Oct 2018 HUG was focused on utilizing the recently released Hashicorp Consul Helm Chart, enabling native integration between Consul and Kubernetes.  Once deployed, the sync service will start querying Kubernetes and Consul for services definitions, and begin the process of adding them to the respective service.

In our example we showcased deploying a new application in Kubernetes (GKE in this case) that is NOT natively Consul aware.  Once the service is present in Kubernetes, the sync service adds the definition to Consul for us.  This effectively allows applications to perform service discovery in either Kubernetes natively or via Consul directly.  This use case in particular ensures that regardless of the service discovery mechanism chosen by app teams, applications can find each other accordingly.

## Requirements

In order to replicate the testing, sample code is provided here. In order to ease the deployment Helm is utilized to deploy the test app, through orchestration via Terraform. This requires Helm be installed in your cluster with the appropriate permissions to deploy applications and services.

**All Consul components were setup using the newly released Hashicorp Consul Helm Chart, and orchestrated with Terraform**  A submodule is included in this repository for reference.

Once you have populated a kubectl config with your cluster credentials, initialize Terraform and apply to deploy the test app.

## Vault HA on Kubernetes

Additionally the demonstration included a reference deployment of Vault in HA mode on Kubernetes. This was built using a Helm Chart with a statefulset along with service definitions. Sample Terraform and Helm code is included in this code repo as well.