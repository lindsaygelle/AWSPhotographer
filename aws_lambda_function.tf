resource "aws_lambda_function" "s3_object_notification" {
  architectures           = ["x86_64"]
  code_signing_config_arn = null
  description             = null
  environment {
    variables = {
      application = var.application
      region      = var.region
    }
  }
  handler       = "main"
  filename      = "./src/lambda_function/s3_object_notification/lambda.zip"
  function_name = "${var.application}S3ObjectNotification"
  layers        = null
  memory_size   = 128
  package_type  = "Zip"
  publish       = false
  runtime       = "provided.al2"
  skip_destroy  = false
  role          = aws_iam_role.lambda_s3_bucket_notification.arn
  timeout       = 3
  tracing_config {
    mode = "Active"
  }
}
