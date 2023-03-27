{
  FunctionName: '{{ tfstate `aws_lambda_function.server.function_name` }}',
  MemorySize: 512,
  Architectures: [
    'x86_64'
  ],
  Role: '{{ tfstate `aws_iam_role.server.arn` }}',
  PackageType: 'Image',
  Code: {
    ImageUri: '{{ tfstate `aws_lambda_function.server.image_uri` }}'
  }
}
