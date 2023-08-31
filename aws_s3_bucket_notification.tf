resource "aws_s3_bucket_notification" "object_created_images_uploaded" {
  bucket = aws_s3_bucket.main.id
  topic {
    events = [
      "s3:ObjectCreated:*"
    ]
    filter_prefix = aws_s3_object.images_uploaded.key
    topic_arn     = aws_sns_topic.s3_object_created_images_uploaded.arn
  }
}
