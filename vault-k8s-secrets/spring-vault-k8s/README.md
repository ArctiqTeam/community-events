# Create the approle for spring-vault-weirdscience

`vault write auth/approle/role/spring-vault-weirdscience policies=spring-vault`
`vault read auth/approle/role/spring-vault-weirdscience/role-id`
`vault write -f auth/approle/role/spring-vault-weirdscience/secret-id`

# Get secret for app (openshift or gke)

`vault kv get secret/weirdscience/openshift`

# Create a secret called secret-id with the retrieved password

```
apiVersion: v1
kind: Secret
metadata:
  name: secret-id

```


`kubectl apply -f secret.yaml`
`oc create -f secret.yaml`



# Create the pod

`kubectl apply -f pod.yaml`
`oc create -f pod.yaml`

# Maven profile

Change the maven profile by changing the value of MVN_PROFILE env var to gke or openshift
