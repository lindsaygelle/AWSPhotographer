resource "aws_iam_role_policy_attachment" "lambda_rekognition_lambda_basic_execution_role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda_rekognition.id
}

resource "aws_iam_role_policy_attachment" "lambda_s3_bucket_notification_lambda_basic_execution_role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda_s3_bucket_notification.id
}

resource "aws_iam_role_policy_attachment" "lambda_s3_bucket_notification_s3_object_read_only_access" {
  policy_arn = aws_iam_policy.s3_object_read_only_access.arn
  role       = aws_iam_role.lambda_s3_bucket_notification.id
}

resource "aws_iam_role_policy_attachment" "lambda_s3_bucket_notification_s3_object_write_only_access" {
  policy_arn = aws_iam_policy.s3_object_write_only_access.arn
  role       = aws_iam_role.lambda_s3_bucket_notification.id
}
