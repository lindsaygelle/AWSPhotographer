resource "aws_lambda_provisioned_concurrency_config" "main" {
  function_name                     = aws_lambda_alias.main.function_name
  provisioned_concurrent_executions = var.provisioned_concurrent_executions
  qualifier                         = aws_lambda_alias.main.function_version
}
