data "http" "github_oidc_configuration" {
  url = "https://token.actions.githubusercontent.com/.well-known/openid-configuration"
}

data "tls_certificate" "github_oidc" {
  url = jsondecode(data.http.github_oidc_configuration.response_body).jwks_uri
}

data "aws_iam_openid_connect_provider" "github_oidc" {
  url = "https://token.actions.githubusercontent.com"
}

data "aws_iam_policy_document" "assume_actions_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Federated"
      identifiers = [data.aws_iam_openid_connect_provider.github_oidc.arn]
    }
    actions = [
      "sts:AssumeRoleWithWebIdentity"
    ]
    condition {
      test     = "StringLike"
      variable = "token.actions.githubusercontent.com:sub"
      values = [
        "repo:AgricolaDevJP/agricoladb-server:*"
      ]
    }
  }
}

resource "aws_iam_role" "actions_role" {
  name               = "agricoladb-server-github-actions-role"
  assume_role_policy = data.aws_iam_policy_document.assume_actions_role.json
}

resource "aws_iam_role_policy_attachment" "actions_role_policy_dynamodb" {
  role       = aws_iam_role.actions_role.id
  policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
}

resource "aws_iam_role_policy_attachment" "actions_role_policy_s3" {
  role       = aws_iam_role.actions_role.id
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

resource "aws_iam_role_policy_attachment" "actions_role_policy_lambda" {
  role       = aws_iam_role.actions_role.id
  policy_arn = "arn:aws:iam::aws:policy/AWSLambda_FullAccess"
}

resource "aws_iam_role_policy_attachment" "actions_role_policy_apigateway" {
  role       = aws_iam_role.actions_role.id
  policy_arn = "arn:aws:iam::aws:policy/AmazonAPIGatewayAdministrator"
}

resource "aws_iam_role_policy_attachment" "actions_role_policy_iam" {
  role       = aws_iam_role.actions_role.id
  policy_arn = "arn:aws:iam::aws:policy/IAMFullAccess"
}

resource "aws_iam_role_policy_attachment" "actions_role_policy_logs" {
  role       = aws_iam_role.actions_role.id
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
}

resource "aws_iam_role_policy_attachment" "actions_role_policy_ecr" {
  role       = aws_iam_role.actions_role.id
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"
}
