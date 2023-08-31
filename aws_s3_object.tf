resource "aws_s3_object" "analytics" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  key          = "analytics/"
}

resource "aws_s3_object" "dynamodb" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  key          = "dynamodb/"
}

resource "aws_s3_object" "images" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  key          = "images/"
}

resource "aws_s3_object" "images_compressed" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  key          = "images/compressed/"
}

resource "aws_s3_object" "images_uploaded" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  key          = "images/uploaded/"
}

resource "aws_s3_object" "inventory" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  key          = "inventory/"
}

resource "aws_s3_object" "logs" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  key          = "logs/"
}

resource "aws_s3_object" "rekognition" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  key          = "rekognition/"
}
