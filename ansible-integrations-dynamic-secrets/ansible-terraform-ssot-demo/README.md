# Ansible Terraform SSOT Demo

Demo content containing examples of executing Ansible with Terraform, Terraform with Ansible, manipulating a centralized inventory, and combining the power of Ansible Tower and Terraform Enterprise.

## GCP Setup

### Project

Create the project to be used for this demo:

```bash
gcloud projects create <project_id>
```

### Service Account Credentials

Set your local credentials directory and filename:

```bash
creds_dir="~/.gcp"
key="$creds_dir/account.json"

mkdir "$creds_dir"
chmod 700 "$creds_dir"
```

Create a service account to provision GCE resources:

```bash
gcloud iam service-accounts create gce-provisioner --project <project_id>
```

Create and download the service account **private** key:

```bash
gcloud iam service-accounts keys create "$key" --iam-account gce-provisioner@<project_id>.iam.gserviceaccount.com
```

Assign policy binding to provision GCE resources:

```bash
gcloud projects add-iam-policy-binding <project_id> --member serviceAccount:gce-provisioner@<project_id>.iam.gserviceaccount.com --role roles/compute.admin
```

### API Services

Enable the GCE service in your project:

```bash
gcloud services enable compute.googleapis.com --project <project_id>
```

## Terraform

Initialize Terraform:

```bash
terraform init
```

Apply Terraform:

```bash
terraform apply
```

## Ansible

### Dependencies

Install the required `google-auth` Python library. In Fedora:

```bash
sudo dnf install python3-google-auth
```
