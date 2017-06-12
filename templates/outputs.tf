output "aws_instance_name" {
  value = "${aws_instance.foo.tags.Name}"
}
