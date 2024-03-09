data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "server" {
  name               = "${var.name}-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy" "lambda_basic_execution" {
  arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "this" {
  role       = aws_iam_role.server.name
  policy_arn = data.aws_iam_policy.lambda_basic_execution.arn
}

resource "aws_lambda_function" "server" {
  function_name = var.name
  package_type  = "Zip"
  role          = aws_iam_role.server.arn
  architectures = [var.lambda_architecture]
  filename      = "./agricoladb-server-lambda.zip"
  handler       = "main"
  memory_size   = 256
  timeout       = 7
  runtime       = "provided.al2023"

  environment {
    variables = {
      ALLOWED_ORIGINS = join(",", var.allowed_origins)
    }
  }

  lifecycle {
    ignore_changes = [
      architectures,
      memory_size,
      timeout,
    ]
  }

  depends_on = [
    terraform_data.server_lambda_archive
  ]
}

resource "aws_lambda_permission" "api" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.server.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.api.execution_arn}/*/*"
}
