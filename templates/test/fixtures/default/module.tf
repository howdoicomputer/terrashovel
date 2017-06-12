provider "aws" {
  region = "us-east-1"
}

resource "random_id" "server" {
  byte_length = 8
}

module "default" {
  name    = "doc_holiday_${random_id.server.hex}"
  source  = "../../../"
  testing = "true"
}

