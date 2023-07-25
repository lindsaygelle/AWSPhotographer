resource "aws_sns_topic" "image" {
  name = "${var.application}-image"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}

resource "aws_sns_topic" "rekognition" {
  name = "${var.application}-rekognition"
  tags = {
    "Account"     = "${var.account}"
    "Application" = "${var.application}"
    "Region"      = "${var.region}"
  }
}
