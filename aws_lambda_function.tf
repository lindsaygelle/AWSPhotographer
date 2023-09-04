resource "aws_lambda_function" "rekognition_detect_faces" {
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
  filename      = "./src/lambda_function/rekognition/detect_faces/lambda.zip"
  function_name = "${var.application}RekognitionDetectFaces"
  layers        = null
  memory_size   = 128
  package_type  = "Zip"
  publish       = false
  runtime       = "go1.x"
  skip_destroy  = false
  role          = aws_iam_role.lambda_rekognition.arn
  timeout       = 3
  tracing_config {
    mode = "Active"
  }
}
