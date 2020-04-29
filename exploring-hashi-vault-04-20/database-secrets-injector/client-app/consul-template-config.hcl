vault {
  renew_token = false

  retry {
    backoff     = "250ms"
    attempts    = 6
    max_backoff = "1m"
  }

  ssl {
    enabled = true
    ca_cert = "/etc/ssl/certs/ca.pem"
  }
}

template {
  destination = "/var/tmp/sql/.pgpass"
  perms       = 0600

  contents = <<EOH
{{- with secret "database/creds/readonly" }}
postgres.weirdscience.labs.arctiq.ca:5432:postgresdb:{{ .Data.username }}:{{ .Data.password }}
{{ end }}
EOH
}

template {
  destination = "/var/tmp/sql/username"
  perms       = 0600

  contents = <<EOH
{{- with secret "database/creds/readonly" }}
{{ .Data.username }}
{{ end }}
EOH
}
