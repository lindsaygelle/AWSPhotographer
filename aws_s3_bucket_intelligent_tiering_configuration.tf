resource "aws_s3_bucket_intelligent_tiering_configuration" "main" {
  bucket = aws_s3_bucket.main.id
  name   = aws_s3_bucket.main.bucket
  tiering {
    access_tier = "ARCHIVE_ACCESS"
    days        = 125
  }
  tiering {
    access_tier = "DEEP_ARCHIVE_ACCESS"
    days        = 180
  }
}
