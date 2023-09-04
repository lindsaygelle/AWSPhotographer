resource "aws_lambda_alias" "rekognition_detect_faces" {
  function_name    = aws_lambda_function.rekognition_detect_faces.function_name
  function_version = "$LATEST"
  name             = aws_lambda_function.rekognition_detect_faces.function_name
}
