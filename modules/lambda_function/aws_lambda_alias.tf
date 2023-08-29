
resource "aws_lambda_alias" "main" {
  function_name    = aws_lambda_function.main.function_name
  function_version = "$LATEST"
  name             = aws_lambda_function.main.function_name
}
