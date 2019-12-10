provider "google" {
  region  = "northamerica-northeast1"
  project = "iceberg-234522"
}

terraform {
  backend "gcs" {
    bucket = "iceberg-actions-bucket"
    prefix = "terraform/state"
  }
}

resource "google_compute_instance" "default" {
  count = tonumber(chomp(file("issues/latest_issue")))
  name         = "assistant-instance-${count.index}"
  machine_type = "f1-micro"
  zone         = "us-east1-b"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
 }
