resource "aws_s3_bucket_notification" "main" {
  bucket = aws_s3_bucket.main.id
  depends_on = [
    aws_lambda_function.s3_object_notification_object_created_image,
    aws_lambda_function.s3_object_notification_object_created_image_compressed
  ]
  /*
  topic {
    events = [
      "s3:ObjectCreated:*"
    ]
    filter_prefix = aws_s3_object.images_uploaded.key
    topic_arn     = aws_sns_topic.s3_object_created_images_uploaded.arn
  }
  */
  lambda_function {
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = aws_s3_object.images_uploaded.key
    filter_suffix       = ".JPG"
    lambda_function_arn = aws_lambda_function.s3_object_notification_object_created_image.arn
  }

  lambda_function {
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = aws_s3_object.images_compressed.key
    filter_suffix       = ".JPG"
    lambda_function_arn = aws_lambda_function.s3_object_notification_object_created_image_compressed.arn
  }
}
