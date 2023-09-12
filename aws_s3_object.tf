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
  depends_on   = [aws_s3_object.images]
  key          = "${aws_s3_object.images.key}compressed/"
}

resource "aws_s3_object" "images_exif" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  depends_on   = [aws_s3_object.images]
  key          = "${aws_s3_object.images.key}exif/"
}

resource "aws_s3_object" "images_uploaded" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  depends_on   = [aws_s3_object.images]
  key          = "${aws_s3_object.images.key}uploaded/"
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

resource "aws_s3_object" "rekognition_detect_faces" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  depends_on   = [aws_s3_object.rekognition]
  key          = "${aws_s3_object.rekognition.key}detect_faces/"
}

resource "aws_s3_object" "rekognition_detect_labels" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  depends_on   = [aws_s3_object.rekognition]
  key          = "${aws_s3_object.rekognition.key}detect_labels/"
}

resource "aws_s3_object" "rekognition_detect_moderation_labels" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  depends_on   = [aws_s3_object.rekognition]
  key          = "${aws_s3_object.rekognition.key}detect_moderation_labels/"
}

resource "aws_s3_object" "rekognition_detect_text" {
  bucket       = aws_s3_bucket.main.id
  content_type = "application/x-directory"
  depends_on   = [aws_s3_object.rekognition]
  key          = "${aws_s3_object.rekognition.key}detect_text/"
}
