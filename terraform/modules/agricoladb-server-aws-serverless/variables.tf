variable "name" {
  type    = string
  default = "agricoladb-server"
}

variable "custom_domain_name" {
  type    = string
  default = null
}

variable "allowed_origins" {
  type    = list(string)
  default = ["http://localhost:4321"]
}

variable "api_throttling_burst_limit" {
  type    = number
  default = 100
}

variable "api_throttling_rate_limit" {
  type    = number
  default = 30
}

variable "api_logs_retention_in_days" {
  type    = number
  default = 7
}

variable "lambda_architecture" {
  type    = string
  default = "arm64"
}
