provider "vault" {
  address = "https://vault.arctiq.ca"
}

data "vault_generic_secret" "gcp_auth" {
  path = "gcp/key/arctiqtim_dns"
}

provider "google" {
  credentials = "${base64decode(data.vault_generic_secret.gcp_auth.data.private_key_data)}"
}

resource "google_dns_record_set" "darkhelmet" {
  name         = "*.arctiqtim.team.arctiq.ca."
  type         = "A"
  ttl          = 300
  managed_zone = "arctiqtim-zone"
  rrdatas      = ["35.203.117.96"]
  project      = "arctiqtim-zone"
}