provider "azurerm" {
}

resource "azurerm_resource_group" "rg" {
  name     = var.rg_name
  location = var.location
  tags     = { 
    project = var.object-tag
  }
}

resource "azurerm_dns_zone" "dns" {
  name                = var.dns_name
  resource_group_name = "${azurerm_resource_group.rg.name}"
  tags                = { 
    project = var.object-tag
  }
}