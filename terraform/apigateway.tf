resource "aws_apigatewayv2_api" "api" {
  name          = "agricoladb-api"
  protocol_type = "HTTP"
  cors_configuration {
    allow_credentials = true
    allow_methods = ["GET", "POST", "OPTIONS"]
    allow_origins = ["http://localhost:4321", "https://db.agricolajp.dev"]
  }
}

resource "aws_apigatewayv2_integration" "server" {
  api_id                 = aws_apigatewayv2_api.api.id
  integration_type       = "AWS_PROXY"
  integration_uri        = aws_lambda_function.server.invoke_arn
  integration_method     = "POST"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "proxy" {
  api_id    = aws_apigatewayv2_api.api.id
  route_key = "ANY /{proxy+}"
  target    = "integrations/${aws_apigatewayv2_integration.server.id}"
}

resource "aws_cloudwatch_log_group" "api" {
  name              = "/aws/apigateway/agricoladb-api"
  retention_in_days = 7
}

resource "aws_apigatewayv2_stage" "default" {
  api_id      = aws_apigatewayv2_api.api.id
  name        = "$default"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api.arn
    format          = jsonencode({ "requestId" : "$context.requestId", "ip" : "$context.identity.sourceIp", "requestTime" : "$context.requestTime", "httpMethod" : "$context.httpMethod", "routeKey" : "$context.routeKey", "status" : "$context.status", "protocol" : "$context.protocol", "responseLength" : "$context.responseLength" })
  }

  default_route_settings {
    throttling_burst_limit = 100
    throttling_rate_limit  = 30
  }
}


data "aws_acm_certificate" "api" {
  domain = "api.db.agricolajp.dev"
}

resource "aws_apigatewayv2_domain_name" "api" {
  domain_name = "api.db.agricolajp.dev"

  domain_name_configuration {
    certificate_arn = data.aws_acm_certificate.api.arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

data "aws_route53_zone" "agricolajp_dev" {
  name = "agricolajp.dev"
}

resource "aws_route53_record" "api" {
  name    = aws_apigatewayv2_domain_name.api.domain_name
  type    = "A"
  zone_id = data.aws_route53_zone.agricolajp_dev.zone_id

  alias {
    name                   = aws_apigatewayv2_domain_name.api.domain_name_configuration[0].target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.api.domain_name_configuration[0].hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_apigatewayv2_api_mapping" "api" {
  api_id      = aws_apigatewayv2_api.api.id
  domain_name = aws_apigatewayv2_domain_name.api.domain_name
  stage       = aws_apigatewayv2_stage.default.name
}
