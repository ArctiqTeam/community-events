variable "AWS_REGION" {
  default = "us-east-1"
}
variable "AMIS" {
  type = "map"
  default = {
    us-east-1 = "ami-82cfb894"
    us-west-2 = "ami-e9503589"
    eu-west-1 = "ami-051b1563"
  }
}
variable "root_volume_size" {
  description = "How big to make the disk in GB"
  default = 12
}
variable "volume_delete_on_termination" {
  description = "Delete the disk volume on VM termination"
  default = true
}
