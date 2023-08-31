data "aws_iam_policy_document" "s3_bucket_notification_s3_object_created_images_uploaded" {
  statement {
    actions = [
      "SNS:Publish"
    ]
    effect = "Allow"
    principals {
      identifiers = [
        "s3.amazonaws.com"
      ]
      type = "Service"
    }
    resources = [
      aws_sns_topic.s3_object_created_images_uploaded.arn
    ]
  }
}
