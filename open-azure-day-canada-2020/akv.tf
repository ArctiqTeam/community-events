# Provision Azure Key Vault
resource "azurerm_key_vault" "vault" {
  name                        = "vault-${random_pet.prefix.id}"
  location                    = azurerm_resource_group.default.location
  resource_group_name         = azurerm_resource_group.default.name
  enabled_for_deployment      = true
  enabled_for_disk_encryption = true
  tenant_id                   = var.tenant_id

  sku_name = "standard"

  tags = {
    environment = "Demo"
  }

  access_policy {
    tenant_id = var.tenant_id

    object_id = "98004206-0c2b-4a39-99c4-9dad12a519b9"

    key_permissions = [
      "get",
      "list",
      "create",
      "delete",
      "update",
      "wrapKey",
      "unwrapKey",
    ]
  }

  access_policy {
    tenant_id = var.tenant_id

    object_id = "c93418d4-00b8-4a52-b9e0-2347c521181f"

    key_permissions = [
      "get",
      "list",
      "create",
      "delete",
      "update",
      "wrapKey",
      "unwrapKey",
    ]
  }

  access_policy {
    tenant_id = var.tenant_id

    object_id = "2c47b890-7aed-4a73-b468-cf4f15cb677f"

    key_permissions = [
      "get",
      "list",
      "create",
      "delete",
      "update",
      "wrapKey",
      "unwrapKey",
    ]
  }

  network_acls {
    default_action = "Allow"
    bypass         = "AzureServices"
  }
}

resource "azurerm_key_vault_key" "generated" {
  name         = var.key_name
  key_vault_id = azurerm_key_vault.vault.id
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "decrypt",
    "encrypt",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey",
  ]
}

output "key_vault_name" {
  value = azurerm_key_vault.vault.name
}
