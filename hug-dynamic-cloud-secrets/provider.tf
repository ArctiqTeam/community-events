provider "vault" {}

data "vault_aws_access_credentials" "creds" {
  backend = "aws"
  role    = "provisioner"
}

provider "aws" {
  access_key = "${data.vault_aws_access_credentials.creds.access_key}"
  secret_key = "${data.vault_aws_access_credentials.creds.secret_key}"

  #   region     = "${data.external.region.result["region"]}"
  region = "${var.AWS_REGION}"
}
