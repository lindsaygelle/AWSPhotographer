resource "aws_s3_bucket" "main" {
  bucket = "${var.account}-${var.application}"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}
