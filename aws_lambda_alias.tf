resource "aws_lambda_alias" "s3_object_notification" {
  function_name    = aws_lambda_function.s3_object_notification.function_name
  function_version = "$LATEST"
  name             = aws_lambda_function.s3_object_notification.function_name
}
