provider "vault" {
  address = "https://<your vault address here>"
}

data "vault_generic_secret" "gcp_auth" {
  path = "<path to your secret key>"
}

provider "google" {
  credentials = "${base64decode(data.vault_generic_secret.gcp_auth.data.private_key_data)}"
}

resource "google_dns_record_set" "<name the record>" {
  name         = "<your DNS name (FQDN) for the node>."
  type         = "A"
  ttl          = 300
  managed_zone = "<a dns zone in gcp>"
  rrdatas      = ["IP for the tower node"]
  project      = "<gcp project providing the zone>"
}


