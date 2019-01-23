provider "vault" {
  address = "https://vault.arctiq.ca"
}

data "vault_generic_secret" "gcp_auth" {
  path = "gcp/key/funixz_dns"
}

provider "google" {
  credentials = "${base64decode(data.vault_generic_secret.gcp_auth.data.private_key_data)}"
}

resource "google_dns_record_set" "towaz" {
  name         = "towaz.funixz.team.arctiq.ca."
  type         = "A"
  ttl          = 300
  managed_zone = "funixz-zone"
  rrdatas      = ["13.66.173.12"]
  project      = "funixz-zone"
}


