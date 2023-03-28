terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }

  backend "s3" {
    bucket = "agricoladevjp-terraform-states"
    region = "ap-northeast-1"
    key    = "AgricolaDevJP/agricoladb-server/terraform_states.tfstate"
  }
}

provider "aws" {
  region = "ap-northeast-1"
}
