data "aws_iam_policy_document" "assume_role_lambda" {
  statement {
    actions = [
      "sts:AssumeRole"
    ]
    effect = "Allow"
    principals {
      identifiers = [
        "lambda.amazonaws.com"
      ]
      type = "Service"
    }
  }
}

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
