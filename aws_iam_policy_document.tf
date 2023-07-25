data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    actions = [
      "sts:AssumeRole"
    ]
    principals {
      identifiers = [
        "lambda.amazonaws.com"
      ]
      type = "Service"
    }
  }
}

data "aws_iam_policy_document" "s3_full_access" {
  statement {
    actions = [
      "s3:*"
    ]
    effect = "Allow"
    resources = [
      "${aws_s3_bucket.main.arn}",
      "${aws_s3_bucket.main.arn}/*"
    ]
  }
}

data "aws_iam_policy_document" "s3_read_only_access" {
  statement {
    actions = [
      "s3:Get*",
      "s3:List*",
      "s3-object-lambda:Get*",
      "s3-object-lambda:List*"
    ]
    effect = "Allow"
    resources = [
      "${aws_s3_bucket.main.arn}",
      "${aws_s3_bucket.main.arn}/*"
    ]
  }
}

data "aws_iam_policy_document" "s3_read_only_object_access" {
  statement {
    actions = [
      "s3:GetObject*"
    ]
    effect = "Allow"
    resources = [
      "${aws_s3_bucket.main.arn}",
      "${aws_s3_bucket.main.arn}/*"
    ]
  }
}

data "aws_iam_policy_document" "s3_write_only_access" {
  statement {
    actions = [
      "s3:Put*"
    ]
    effect = "Allow"
    resources = [
      "${aws_s3_bucket.main.arn}",
      "${aws_s3_bucket.main.arn}/*"
    ]
  }
}

data "aws_iam_policy_document" "s3_write_only_object_access" {
  statement {
    actions = [
      "s3:PutObject*"
    ]
    effect = "Allow"
    resources = [
      "${aws_s3_bucket.main.arn}",
      "${aws_s3_bucket.main.arn}/*"
    ]
  }
}

data "aws_iam_policy_document" "sns_subscribe_only_access" {
  statement {
    actions = [
      "sns:Subscribe",
      "sns:Unsubscribe"
    ]
    effect = "Allow"
    resources = [
      aws_sns_topic.image.arn,
      aws_sns_topic.rekognition.arn
    ]
  }
}

data "aws_iam_policy_document" "sns_subscribe_image_only_access" {
  statement {
    actions = [
      "sns:Subscribe",
      "sns:Unsubscribe"
    ]
    effect = "Allow"
    resources = [
      aws_sns_topic.image.arn
    ]
  }
}

data "aws_iam_policy_document" "sns_subscribe_rekognition_only_access" {
  statement {
    actions = [
      "sns:Subscribe",
      "sns:Unsubscribe"
    ]
    effect = "Allow"
    resources = [
      aws_sns_topic.rekognition.arn
    ]
  }
}

data "aws_iam_policy_document" "sns_topic_policy_image" {
  statement {
    actions = [
      "SNS:Publish"
    ]
    condition {
      test = "ArnLike"
      values = [
        aws_s3_bucket.main.arn
      ]
      variable = "aws:SourceArn"
    }
    effect = "Allow"
    principals {
      identifiers = [
        "s3.amazonaws.com"
      ]
      type = "Service"
    }
    resources = [
      aws_sns_topic.image.arn
    ]
  }
}
