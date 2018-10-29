resource "helm_release" "vault" {
  name          = "${var.vault_name}"
  namespace     = "${var.vault_name}"
  chart         = "${var.vault_name}/"
  recreate_pods = true

  values = [
    "name: ${var.vault_name}",
    "image: vault",
    "image_version: ${var.vault_image_version}",
  ]
}

module "vault-init" {
  source     = "github.com/arctiqtim/terraform-shell-resource"
  command    = "kubectl exec ${var.vault_name}-0 -n ${var.vault_name} -it -- vault operator init -format json"
  depends_id = "${helm_release.vault.id}"
}
