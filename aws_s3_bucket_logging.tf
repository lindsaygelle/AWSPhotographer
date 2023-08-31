resource "aws_s3_bucket_logging" "main" {
  bucket        = aws_s3_bucket.main.id
  target_bucket = aws_s3_bucket.main.id
  target_prefix = aws_s3_object.logs.key
}
