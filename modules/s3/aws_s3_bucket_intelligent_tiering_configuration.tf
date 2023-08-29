resource "aws_s3_bucket_intelligent_tiering_configuration" "main" {
  bucket = aws_s3_bucket.main.id
  name   = "bucket"
  tiering {
    access_tier = "ARCHIVE_ACCESS"
    days        = var.intelligent_tiering_configuration.access_tier.days
  }
  tiering {
    access_tier = "DEEP_ARCHIVE_ACCESS"
    days        = var.intelligent_tiering_configuration.deep_archive_access.days
  }
}
