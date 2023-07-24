resource "aws_iam_role" "lambda" {
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
  description        = "This role will grant permissions required to use Lambda."
  path               = "/${var.application}/"
  name               = "LambaRole"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}

resource "aws_iam_role" "lambda_image" {
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
  description        = "This role will grant permissions required for Lambda to process SNS topic ${aws_sns_topic.image.arn}."
  path               = "/${var.application}/"
  name               = "LambdaSNSImageRole"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}


resource "aws_iam_role" "lambda_rekognition" {
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
  description        = "This role will grant permissions required for Lambda to process SNS topic ${aws_sns_topic.rekognition.arn}."
  path               = "/${var.application}/"
  name               = "LambdaSNSRekognitionRole"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}
