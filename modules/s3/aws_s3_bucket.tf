resource "aws_s3_bucket" "main" {
  bucket              = var.bucket_name
  bucket_prefix       = var.bucket_prefix
  force_destroy       = true
  object_lock_enabled = false
  tags                = var.tags
}
