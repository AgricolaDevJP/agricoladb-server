terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }

  backend "s3" {
    bucket                  = "agricoladevjp-terraform-states"
    region                  = "ap-northeast-1"
    key                     = "AgricolaDevJP/agricoladb-server/terraform_states.tfstate"
    shared_credentials_file = "~/.aws/credentials"
    profile                 = "AgricolaDevJP-admin"
  }
}

provider "aws" {
  region  = "ap-northeast-1"
  profile = "AgricolaDevJP-admin"
}
