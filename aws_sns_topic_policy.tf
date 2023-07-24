resource "aws_sns_topic_policy" "image" {
  arn    = aws_sns_topic.image.arn
  policy = data.aws_iam_policy_document.sns_topic_policy_image.json
}
