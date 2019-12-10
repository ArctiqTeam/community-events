provider "google" {
  region  = "us-central1"
  project = "blizzard-253119"
}

terraform {
  backend "gcs" {
    bucket = "terraform-back3nd"
    prefix = "terraform/state"
  }
}

resource "google_storage_bucket_object" "index" {
  name   = "index.html"
  source = "pages/index.html"
  bucket = "www.sourpatch.net"
}

resource "google_storage_bucket_object" "error" {
  name   = "error.html"
  source = "pages/error.html"
  bucket = "www.sourpatch.net"
}
