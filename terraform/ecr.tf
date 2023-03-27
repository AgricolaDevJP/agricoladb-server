locals {
  ecr_lifecycle_policy = {
    rules = [
      {
        action = {
          type = "expire"
        }
        rulePriority = 1
        selection = {
          tagStatus   = "untagged"
          countType   = "sinceImagePushed"
          countUnit   = "days"
          countNumber = 1
        }
      }
    ]
  }
}

resource "aws_ecr_repository" "server_lambda" {
  name                 = "agricoladb-server-lambda"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_lifecycle_policy" "server_lambda" {
  repository = aws_ecr_repository.server_lambda.name
  policy     = jsonencode(local.ecr_lifecycle_policy)
}
