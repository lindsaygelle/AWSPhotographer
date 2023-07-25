resource "aws_s3_bucket_notification" "main" {
  bucket = aws_s3_bucket.main.id
  topic {
    events        = ["s3:ObjectCreated:*"]
    filter_suffix = ".jpeg"
    topic_arn     = aws_sns_topic.image.arn
  }
  topic {
    events        = ["s3:ObjectCreated:*"]
    filter_suffix = ".png"
    topic_arn     = aws_sns_topic.image.arn
  }
  topic {
    events        = ["s3:ObjectCreated:*"]
    filter_suffix = ".raw"
    topic_arn     = aws_sns_topic.image.arn
  }
}
