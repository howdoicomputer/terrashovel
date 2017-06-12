# About

Terrashovel is a very small CLI that scaffolds out a Terraform module.

# Installation

If you have Go installed and setup:

`go get -u github.com/howdoicomputer/terrashovel`

Otherwise, grab the [latest release](https://github.com/howdoicomputer/terrashovel/releases) and put it somewhere in your $PATH.

# Usage

```
terrashovel create tf_test_module

cd tf_test_module
bundle install
bundle exec kitchen test
```

The scaffolding for the generated module deploys a single t2.micro instance into AWS and includes a single test to verify that the instance exists. The `main.tf` for the module looks like this:

```
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
```

In a future release, I'll be moving all *working* code into files that are only generated under a `--generate-example` flag and having the default scaffolding generate non-functional, empty files.

# Assumptions

The test suite used to validate created infrastructure is written in Ruby. While not a requirement for Terrashovel per se, a Ruby environment is required to install and run the suite (`bundle exec kitchen test`). To satisfy that requirement, I recommend [RVM](https://rvm.io/).

---
