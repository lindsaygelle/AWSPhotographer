resource "aws_lambda_permission" "s3_object_notification_object_created_image" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.s3_object_notification_object_created_image.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.main.arn
  statement_id  = "AllowS3Invoke"
}
