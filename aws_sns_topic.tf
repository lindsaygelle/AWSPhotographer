resource "aws_sns_topic" "s3_object_images_uploaded" {
  name = "${lower(var.application)}-${replace(trimsuffix(aws_s3_object.images_uploaded.key, "/"), "/", "-")}"
}
