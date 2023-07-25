terraform {
  backend "s3" {
    bucket = "385739365063"
    key    = "AWS/Photographer/terraform.tfstate"
    region = "us-east-1"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.4.0"
    }
  }
}
