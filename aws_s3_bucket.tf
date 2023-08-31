resource "aws_s3_bucket" "main" {
  bucket = "${data.aws_caller_identity.main.account_id}-${lower(var.application)}"
}
