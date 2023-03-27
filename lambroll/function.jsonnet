{
  FunctionName: '{{ tfstate `aws_lambda_function.server.function_name` }}',
  MemorySize: 256,
  Architectures: [
    'arm64'
  ],
  Timeout: 7,
  Role: '{{ tfstate `aws_iam_role.server.arn` }}',
  PackageType: 'Image',
  Code: {
    ImageUri: '{{ tfstate `aws_lambda_function.server.image_uri` }}'
  },
}
