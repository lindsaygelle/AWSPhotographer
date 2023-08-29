resource "aws_s3_bucket_metric" "main" {
  bucket = aws_s3_bucket.main.id
  name   = "bucket"
}
