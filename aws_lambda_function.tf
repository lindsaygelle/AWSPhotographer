resource "aws_lambda_function" "s3_object_notification_object_created_image" {
  architectures           = ["x86_64"]
  code_signing_config_arn = null
  description             = null
  environment {
    variables = {
      APPLICATION                        = var.application
      REGION                             = var.region
      S3_BUCKET_FOLDER_IMAGES_COMPRESSED = aws_s3_object.images_compressed.key
      S3_BUCKET_FOLDER_IMAGES_EXIF       = aws_s3_object.images_exif.key
      S3_BUCKET_FOLDER_IMAGES_UPLOADED   = aws_s3_object.images_uploaded.key
    }
  }
  handler          = "main"
  filename         = "./src/lambda_function/s3_object_notification/object_created/image/lambda.zip"
  function_name    = "${var.application}S3ObjectNotificationObjectCreatedImageUploaded"
  layers           = null
  memory_size      = 128
  package_type     = "Zip"
  publish          = false
  runtime          = "provided.al2"
  skip_destroy     = false
  source_code_hash = sha256("./src/lambda_function/s3_object_notification/object_created/image/lambda.zip")
  role             = aws_iam_role.lambda_s3_bucket_notification.arn
  timeout          = 60
  tracing_config {
    mode = "Active"
  }
}

resource "aws_lambda_function" "s3_object_notification_object_created_image_compressed" {
  architectures           = ["x86_64"]
  code_signing_config_arn = null
  description             = null
  environment {
    variables = {
      APPLICATION                                           = var.application
      REGION                                                = var.region
      S3_BUCKET_FOLDER_IMAGES_COMPRESSED                    = aws_s3_object.images_compressed.key
      S3_BUCKET_FOLDER_REKOGNITION_DETECT_FACES             = aws_s3_object.rekognition_detect_faces.key
      S3_BUCKET_FOLDER_REKOGNITION_DETECT_LABELS            = aws_s3_object.rekognition_detect_labels.key
      S3_BUCKET_FOLDER_REKOGNITION_DETECT_MODERATION_LABELS = aws_s3_object.rekognition_detect_moderation_labels.key
      S3_BUCKET_FOLDER_REKOGNITION_DETECT_TEXT              = aws_s3_object.rekognition_detect_text.key
    }
  }
  handler          = "main"
  filename         = "./src/lambda_function/s3_object_notification/object_created/image_compressed/lambda.zip"
  function_name    = "${var.application}S3ObjectNotificationObjectCreatedImageCompressed"
  layers           = null
  memory_size      = 128
  package_type     = "Zip"
  publish          = false
  runtime          = "provided.al2"
  skip_destroy     = false
  source_code_hash = sha256("./src/lambda_function/s3_object_notification/object_created/image_compressed/lambda.zip")
  role             = aws_iam_role.lambda_s3_bucket_notification.arn
  timeout          = 60
  tracing_config {
    mode = "Active"
  }
}
