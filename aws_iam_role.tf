resource "aws_iam_role" "lambda_s3_bucket_notification" {
  assume_role_policy = data.aws_iam_policy_document.assume_role_lambda.json
  name               = "LambdaS3BucketNotification"
  path               = "/${var.application}/"
}

resource "aws_iam_role" "lambda_rekognition" {
  assume_role_policy = data.aws_iam_policy_document.assume_role_lambda.json
  name               = "LambdaRekognition"
  path               = "/${var.application}/"
}
