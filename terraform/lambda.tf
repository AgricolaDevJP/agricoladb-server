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

resource "aws_iam_role" "iam_for_lambda" {
  name               = "agricoladb-server-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy" "lambda_basic_execution" {
  arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "lambda_basic_execution" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = data.aws_iam_policy.lambda_basic_execution.arn
}

resource "aws_lambda_function" "server" {
  function_name = "agricoladb-server"
  package_type  = "Image"
  image_uri     = "${aws_ecr_repository.this.repository_url}:latest"
  role          = aws_iam_role.iam_for_lambda.arn
  architectures = ["arm64"]
  memory_size   = 512

  lifecycle {
    ignore_changes = [
      image_uri,
    ]
  }
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.server.function_name
  principal     = "apigateway.amazonaws.com"

  # The /*/* portion grants access from any method on any resource
  # within the API Gateway "REST API".
  source_arn = "${aws_api_gateway_rest_api.this.execution_arn}/*/*"
}
