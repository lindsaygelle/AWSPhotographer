resource "aws_s3_bucket_inventory" "main" {
  bucket                   = aws_s3_bucket.main.id
  included_object_versions = "All"
  destination {
    bucket {
      bucket_arn = aws_s3_bucket.main.arn
      format     = "CSV"
      prefix     = "inventory/bucket/"
    }
  }
  name = "bucket"
  schedule {
    frequency = "Daily"
  }
}
