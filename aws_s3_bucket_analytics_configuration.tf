resource "aws_s3_bucket_analytics_configuration" "main" {
  bucket = aws_s3_bucket.main.id
  name   = aws_s3_bucket.main.bucket
}
