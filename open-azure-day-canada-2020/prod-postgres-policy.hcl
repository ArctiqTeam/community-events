path "database/creds/readonly" {
  capabilities = ["read"]
}

path "/sys/leases/renew" {
  capabilities = ["update"]
}
