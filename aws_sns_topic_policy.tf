resource "aws_sns_topic_policy" "s3_object_created_images_uploaded" {
  arn    = aws_sns_topic.s3_object_created_images_uploaded.arn
  policy = data.aws_iam_policy_document.s3_bucket_notification_s3_object_created_images_uploaded.json
}
