resource "aws_iam_policy" "s3_full_access" {
  description = "Provides full access to the photographer bucket via the AWS Management Console."
  path        = "/${var.application}/"
  policy      = data.aws_iam_policy_document.s3_full_access.json
  name        = "S3FullAccess"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}

resource "aws_iam_policy" "s3_read_only_access" {
  description = "Provides read only access to the photographer bucket via the AWS Management Console."
  path        = "/${var.application}/"
  policy      = data.aws_iam_policy_document.s3_read_only_access.json
  name        = "S3ReadOnlyAccess"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}

resource "aws_iam_policy" "s3_read_only_object_access" {
  description = "Provides read only access to objects in the photographer bucket via the AWS Management Console."
  path        = "/${var.application}/"
  policy      = data.aws_iam_policy_document.s3_read_only_object_access.json
  name        = "S3ReadOnlyObjectAccess"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}

resource "aws_iam_policy" "s3_write_only_access" {
  description = "Provides write only access to the photographer bucket via the AWS Management Console."
  path        = "/${var.application}/"
  policy      = data.aws_iam_policy_document.s3_write_only_access.json
  name        = "S3WriteOnlyAccess"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}

resource "aws_iam_policy" "s3_write_only_object_access" {
  description = "Provides write only access to objects in the photographer bucket via the AWS Management Console."
  path        = "/${var.application}/"
  policy      = data.aws_iam_policy_document.s3_write_only_object_access.json
  name        = "S3WriteOnlyObjectAccess"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}

resource "aws_iam_policy" "sns_subscribe_only_access" {
  description = "Provides subscription only access to SNS topics."
  path        = "/${var.application}/"
  policy      = data.aws_iam_policy_document.sns_subscribe_only_access.json
  name        = "SNSSubscribeOnlyAccess"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}

resource "aws_iam_policy" "sns_subscribe_image_only_access" {
  description = "Provides subscription only access to SNS topic ${aws_sns_topic.image.arn}."
  path        = "/${var.application}/"
  policy      = data.aws_iam_policy_document.sns_subscribe_image_only_access.json
  name        = "SNSSubscribeOnlyImageAccess"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}

resource "aws_iam_policy" "sns_subscribe_rekognition_only_access" {
  description = "Provides subscription only access to SNS topic ${aws_sns_topic.rekognition.arn}."
  path        = "/${var.application}/"
  policy      = data.aws_iam_policy_document.sns_subscribe_rekognition_only_access.json
  name        = "SNSSubscribeOnlyRekognitionAccess"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}
