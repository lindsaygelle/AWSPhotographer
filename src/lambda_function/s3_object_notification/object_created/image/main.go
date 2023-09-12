package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	fileDirectory = "/tmp/"
)

var (
	s3BucketFolderImagesCompressed = os.Getenv("S3_BUCKET_FOLDER_IMAGES_COMPRESSED")
	s3BucketFolderImagesExif       = os.Getenv("S3_BUCKET_FOLDER_IMAGES_EXIF")
	s3BucketFolderImagesUploaded   = os.Getenv("S3_BUCKET_FOLDER_IMAGES_UPLOADED")
)

func handler(context context.Context, event *events.S3Event) {
	log.Printf("S3_BUCKET_FOLDER_IMAGES_COMPRESSED=%s S3_BUCKET_FOLDER_IMAGES_EXIF=%s S3_BUCKET_FOLDER_IMAGES_UPLOADED=%s",
		s3BucketFolderImagesCompressed,
		s3BucketFolderImagesExif,
		s3BucketFolderImagesUploaded)
	if len(s3BucketFolderImagesCompressed) == 0 {
		log.Fatalf("S3_BUCKET_FOLDER_IMAGES_COMPRESSED")
	}
	if len(s3BucketFolderImagesExif) == 0 {
		log.Fatalf("S3_BUCKET_FOLDER_IMAGES_EXIF")
	}
	if len(s3BucketFolderImagesUploaded) == 0 {
		log.Fatalf("S3_BUCKET_FOLDER_IMAGES_UPLOADED")
	}
	if s3BucketFolderImagesUploaded == s3BucketFolderImagesCompressed {
		log.Fatalf("S3_BUCKET_FOLDER_IMAGES_COMPRESSED == S3_BUCKET_FOLDER_IMAGES_UPLOADED")
	}
	session, err := session.NewSession()
	if err != nil {
		log.Fatalln(err)
	}
	processS3Event(session, event)
}

func main() {
	lambda.Start(handler)
}
