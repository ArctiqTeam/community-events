# This is a demo on dynamic secrets in the cloud - Hashicorp User Group May 2018

Overall the purpose of this demonstration is to showcase the ability via Hashicorp Vault to authenticate with GitHub credentials, retrieve a token.  Following the token being issued, you leverage the token to dynamically create an AWS IAM account that only lives for a few minutes, to create an ec2 instance with Terraform.

Once Terraform has done its work (including putting some configuration into the instance re SSH) Ansible take over to configure a service on the instance.  The userdata information is provided here as well for reference for both SSH OTP and SSH CA methods of logging into the host.  For our purposes in the demo we are leveraging SSH CA for the Ansible playbook to authenticate into the instance once provisioned.

All Vault configurations were done according to the standard documentation on the AWS secrets, and SSH CA secrets engines.

## First Generate an access token for use in Vault ... you can skip this if you already have a token

Here is an example with GitHub Auth:

```
curl -k --request POST --data '{"token": "PERSONAL_ACCESS_TOKEN"}' http://<vault_addr>/v1/auth/github/login
{"request_id":"cec4e95b-kith-f29d-2399-036c7dddgdfb","lease_id":"","renewable":false,"lease_duration":0,"data":null,"wrap_info":null,"warnings":null,"auth":{"client_token":"33632a1c-gtef-f1da-fae8-d9d4jtis7fec","accessor":"7ca2281e-cd88-364f-e9fb-f48a0b0fb5d4","policies":["cloud-provisioners","default"],"metadata":{"org":"ArctiqTeam","username":"arctiqtim"},"lease_duration":2764800,"renewable":true,"entity_id":"da5a501e-d61a-6144-4903-68fcd7193d6f"}}
```

Use the `client_token` and set the environment var `$VAULT_TOKEN` with that token

## Now the Terraform config is ready

Ensure you have set the Vault Address variable by:
`export VAULT_ADDR=<your vault server>`

Both the Terraform plan and Ansible playbooks will use the 2 exported vars to connect to vault and retrieve credentials.

You can either execute the 2 separately or together via the `deploy.sh` wrapper script, which will build the ec2 instance with Terraform and then execute the Ansible playbook with the Terraform inventory from state.