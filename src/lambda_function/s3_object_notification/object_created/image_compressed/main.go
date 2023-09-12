package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	s3BucketFolderImagesCompressed       = getEnvironmentVariable("S3_BUCKET_FOLDER_IMAGES_COMPRESSED")
	s3BucketFolderRekognitionDetectFaces = getEnvironmentVariable("S3_BUCKET_FOLDER_REKOGNITION_DETECT_FACES")
)

func createAWSSession() *session.Session {
	awsSession, err := session.NewSession(nil)
	if err != nil {
		log.Fatalf("Session: Error=%s", err)
	}
	return awsSession
}

func getEnvironmentVariable(key string) string {
	environmentValue := os.Getenv(key)
	if len(environmentValue) == 0 {
		log.Fatalf("%s is not set", key)
	}
	return environmentValue
}

func handler(context context.Context, s3Event *events.S3Event) {
	log.Printf("S3_BUCKET_FOLDER_IMAGES_COMPRESSED=%s S3_BUCKET_FOLDER_REKOGNITION_DETECT_FACES=%s",
		s3BucketFolderImagesCompressed,
		s3BucketFolderRekognitionDetectFaces)

	if s3BucketFolderImagesCompressed == s3BucketFolderRekognitionDetectFaces {
		log.Fatalf("S3_BUCKET_FOLDER_IMAGES_COMPRESSED == S3_BUCKET_FOLDER_REKOGNITION")
	}

	awsSession := createAWSSession()
	processS3Event(awsSession, s3Event)
}

func main() {
	lambda.Start(handler)
}
