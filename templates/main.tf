resource "aws_instance", "${var.name}" {
  instance_type = "t2.micro"

  tags {
    Name = "${var.name}"
  }
}
