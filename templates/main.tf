data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "foo" {
  instance_type = "t2.micro"
  ami           = "${data.aws_ami.ubuntu.id}"

  tags {
    Name    = "${var.name}"
    Testing = "${var.testing}"
  }
}

