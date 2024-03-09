terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    github = {
      source = "integrations/github"
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

provider "github" {}
