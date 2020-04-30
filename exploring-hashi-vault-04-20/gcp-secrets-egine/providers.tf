provider "vault" {
  address = "https://vault.marc-leblanc.team.arctiq.ca:8200"
}

data "vault_generic_secret" "gcp_auth" {
  path = "gcp/key/project-factory-roleset"
}

provider "google" {
    credentials = base64decode(data.vault_generic_secret.gcp_auth.data.private_key_data)
}