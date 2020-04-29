```shell
kubectl create serviceaccount vault-auth
kubectl apply -f gke/clusterrolebinding.yaml

export VAULT_SA_NAME=$(kubectl get sa vault-auth -o jsonpath="{.secrets[*]['name']}")
export SA_JWT_TOKEN=$(kubectl get secret $VAULT_SA_NAME -o jsonpath="{.data.token}" | base64 --decode; echo)
export SA_CA_CRT=$(kubectl get secret $VAULT_SA_NAME -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo)
export K8S_HOST=<something>

vault auth enable -path=gke-k8s kubernetes
vault write auth/gke-k8s/config \
        token_reviewer_jwt="$SA_JWT_TOKEN" \
        kubernetes_host="https://$K8S_HOST:443" \
        kubernetes_ca_cert="$SA_CA_CRT"

vault write auth/gke-k8s/role/postgres-client \
        bound_service_account_names=default \
        bound_service_account_namespaces=postgres-client \
        policies=prod-postgres \
        ttl=5m

kubectl create namespace postgres-client

kubectl apply -f postgres/

kubectl port-forward svc/postgres 5432
psql -h localhost -U postgresadmin --password -p 5432 postgresdb < world.sql

psql -h localhost -U postgresadmin --password -p 5432 postgresdb

# Test and do a SELECT * FROM city;

psql -h localhost -U postgresadmin --password -p 5432 postgresdb
CREATE ROLE vault_user WITH SUPERUSER LOGIN CREATEROLE;
\password vault_user;

vault secrets enable database
vault write database/config/prod-postgres \
  plugin_name="postgresql-database-plugin" \
  allowed_roles="readonly" \
  connection_url='postgresql://{{username}}:{{password}}@postgres:5432/postgresdb?sslmode=disable' \
  username=vault_user \
  password='<password>'

vault write database/roles/readonly \
  db_name="prod-postgres" \
  creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
    GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
  default_ttl="5m" \
  max_ttl="12h"

vault policy write prod-postgres gke/prod-postgres-policy.hcl

TEST
vault read database/creds/readonly
```
