resource "aws_iam_policy" "s3_object_read_only_access" {
  path   = "/${var.application}/"
  policy = data.aws_iam_policy_document.s3_object_read_only_access.json
  name   = "${var.application}S3ObjectReadOnlyAccess"
}

resource "aws_iam_policy" "s3_object_write_only_access" {
  path   = "/${var.application}/"
  policy = data.aws_iam_policy_document.s3_object_write_only_access.json
  name   = "${var.application}S3ObjectWriteOnlyAccess"
}
