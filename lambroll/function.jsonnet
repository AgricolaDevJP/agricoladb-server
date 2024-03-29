{
  FunctionName: '{{ tfstate `aws_lambda_function.server.function_name` }}',
  MemorySize: 256,
  Architectures: [
    'arm64',
  ],
  Timeout: 7,
  Role: '{{ tfstate `aws_iam_role.server.arn` }}',
  PackageType: 'Image',
  Code: {
    ImageUri: '{{ tfstate `aws_lambda_function.server.image_uri` }}',
  },
  Environment: {
    Variables: {
      ALLOWED_ORIGINS: 'http://localhost:4321,https://db.agricolajp.dev,https://*.agricoladb-viewer.pages.dev',
      GO_ENV: 'production',
      GO_LOG: 'warn'
    },
  },
}
