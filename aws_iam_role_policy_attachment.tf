resource "aws_iam_role_policy_attachment" "lambda_rekognition_lambda_basic_execution_role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda_rekognition.id
}

resource "aws_iam_role_policy_attachment" "lambda_s3_bucket_notification_lambda_basic_execution_role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda_s3_bucket_notification.id
}
