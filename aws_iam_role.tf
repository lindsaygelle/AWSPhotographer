resource "aws_iam_role" "lambda_s3_bucket_notification" {
  assume_role_policy = data.aws_iam_policy_document.assume_role_lambda.json
  name               = "${var.application}LambdaS3BucketNotification"
  path               = "/${var.application}/"
}
