variable "appId" {
  description = "Azure Kubernetes Service Cluster service principal"
}

variable "password" {
  description = "Azure Kubernetes Service Cluster password"
}

variable "tenant_id" {
  description = "Azure Tenant ID"
}

variable "key_name" {
  description = "Azure Key Vault key name"
  default     = "gen-key"
}
