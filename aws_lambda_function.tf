resource "aws_lambda_function" "rekognition_detect_faces" {
  architectures           = null
  code_signing_config_arn = null
  description             = null
  environment {
    variables = {
      application = var.application
      region      = var.region
    }
  }
  function_name = "${var.application}RekognitionDetectFaces"
  layers        = null
  memory_size   = 128
  package_type  = null
  publish       = false
  runtime       = null
  skip_destroy  = false
  role          = aws_iam_role.lambda_rekognition.arn
  timeout       = 3
  tracing_config {
    mode = "Active"
  }
}
