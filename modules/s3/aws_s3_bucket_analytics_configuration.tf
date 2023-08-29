resource "aws_s3_bucket_analytics_configuration" "main" {
  bucket = aws_s3_bucket.main.id
  name   = "bucket"
  storage_class_analysis {
    data_export {
      destination {
        s3_bucket_destination {
          bucket_arn = aws_s3_bucket.main.arn
          prefix     = "analytics/bucket/"
        }
      }
    }
  }
}
