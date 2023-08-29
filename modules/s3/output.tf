// aws_s3_bucket
output "arn" {
  value = aws_s3_bucket.main.arn
}

output "bucket_name" {
  value = aws_s3_bucket.main.bucket
}

output "id" {
  value = aws_s3_bucket.main.id
}
