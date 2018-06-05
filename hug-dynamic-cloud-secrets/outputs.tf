output "address" {
  value = "${aws_instance.web.public_ip}"
}

output "ssh" {
  value = "ssh ${aws_instance.web.tags.sshUser}@${aws_instance.web.public_ip}"
}
