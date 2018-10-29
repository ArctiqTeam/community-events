provider "kubernetes" {
  host = "https://gke_cluster"
}

provider "helm" {
  install_tiller  = true
  service_account = "helm"
}
