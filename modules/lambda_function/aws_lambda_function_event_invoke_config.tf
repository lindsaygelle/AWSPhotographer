
resource "aws_lambda_function_event_invoke_config" "main" {
  dynamic "destination_config" {
    for_each = !(var.event_invoke_destination_config == null) ? [var.event_invoke_destination_config] : []
    content {
      on_failure {
        destination = event_invoke_destination_config.value.on_failure.destination_arn
      }
      on_success {
        destination = event_invoke_destination_config.value.on_success.destination_arn
      }
    }
  }
  function_name                = aws_lambda_alias.main.function_name
  maximum_event_age_in_seconds = var.maximum_event_age_in_seconds
  maximum_retry_attempts       = var.maximum_retry_attempts
  qualifier                    = aws_lambda_alias.main.function_version
}
