# GCP Secrets Engine and using with Terraform

In the demo for this event, we used a Terraform module called [Project Factory](https://github.com/terraform-google-modules/terraform-google-project-factory) to provision a project in GCP using a dynamic service account key as well as create a Vault roleset that can be used to provision GKE into that new project. Here are some of the code samples used and steps to set up.

## Enable the GCP Secrets Engine

 ```bash
$vault secrets enable --path=gcp gcp
```

## Create a Service Account in GCP for Vault

A Service Account with appropriate IAM roles on the appropriate Project/Folder/Organization is required. 
* Grant `Project IAm Admin` & `Service Account Admin` roles
* Create a key

## Write the key to Vault gcp config
```bash 
$vault write gcp/config credentials=@my-credentials.json
```

## Create the Project Factory Roleset
[project-factory-roleset.hcl](project_factory_roleset.hcl) - Update Folder ID and Org ID
```bash
vault write gcp/roleset/project-factory-roleset \
    project="marc-sandbox-274815" \
    secret_type="access_token" \
    bindings=@project_factory_roleset.hcl
```

## Project Request Terraform

```terraform
module "arc-demo-beer-project" { # Update Module name accordingly
  source  = "tfe.marc-leblanc.team.arctiq.ca/Arctiq/project-factory/google" # Get this from your TFE instance Module path
  version = "7.1.0"
  name = "arc-demo-beer-project" ## Project ID
  org_id = var.org_id  ## Set as Var on the TFE Workspace
  billing_account = var.billing_account ## Set as Var on the TFE Workspace
  folder_id = var.folder_id	 ## Set as Var on the TFE Workspace
  default_service_account = "keep"
  skip_gcloud_download = true
  activate_apis = ["compute.googleapis.com", "container.googleapis.com"] # Change to desired  APIs
}


resource "vault_gcp_secret_roleset" "arc-demo-beer-project-roleset" { # Update roleset-name accordingly
  backend      = "gcp"
  roleset      = "arc-demo-beer-project-roleset" # Update roleset-name accordingly
  secret_type  = "service_account_key"
  project      = "lhq-vaultdemo"    # This is the project where the Service Account will be created. Applying this code will cause the SA to be created, but not a key
  token_scopes = ["https://www.googleapis.com/auth/cloud-platform"]

  binding {
    resource = "//cloudresourcemanager.googleapis.com/projects/arc-demo-beer-project" # what resource are you applying roles to
        roles = ["roles/compute.admin", "roles/container.admin", "roles/iam.serviceAccountUser"] # What roles are associated with this resource
  }
  # Add a dependancy on the module output. This will help with 
  # any race conditions that may be encountered
  depends_on = [module.arc-demo-beer-project.project_id] 
}
```

## Provider Terraform

[provider.tf](./providers.tf)
```terraform
provider "vault" {
  address = "https://vault.marc-leblanc.team.arctiq.ca:8200" ## Address to your Vault instance. Alternative set as an Environment Variable on TFE Workspace
}

data "vault_generic_secret" "gcp_auth" {
  path = "gcp/key/project-factory-roleset" # Path to roleset
}

provider "google" {
    credentials = base64decode(data.vault_generic_secret.gcp_auth.data.private_key_data) # Secret retrieved from Vault
}

```

## TFE SetUp

On each using Vault, create 2 Environment variables
1) VAULT_ADDR = the address to your vault instance
2) VAULT_TOKEN = the Vault access token. **Mark this as sensitive when  creating**

## GKE Demo

Bonus, here is the code to create the GKE clusters as shown in the demo. The only additional step is on the workspaces for these, you need to also define a Terraform variable for the `project_name` where the cluster is being deployed and one for `cluster_name` to be given to the GKE cluster being built. 
[GKE Demo](./gke-demo/)