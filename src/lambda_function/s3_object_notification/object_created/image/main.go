// Package main serves as the entry point for an AWS Lambda function designed to process S3 events.
package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Constants for defining the file directory where temporary files are stored.
const (
	fileDirectory = "/tmp/"
)

// Global variables to store S3 bucket folder names.
var (
	s3BucketFolderImagesCompressed string
	s3BucketFolderImagesExif       string
	s3BucketFolderImagesUploaded   string
)

// createAWSSession creates and returns a new AWS session.
func createAWSSession() *session.Session {
	awsSession, err := session.NewSession(nil)
	if err != nil {
		log.Fatalf("Session: Error=%s", err)
	}
	return awsSession
}

// getEnvironmentVariable retrieves an environment variable by its key and returns its value.
// It exits the program with an error if the variable is not set.
func getEnvironmentVariable(key string) string {
	environmentValue := os.Getenv(key)
	if len(environmentValue) == 0 {
		log.Fatalf("%s is not set", key)
	}
	return environmentValue
}

// handler is the AWS Lambda function that processes S3 events.
func handler(context context.Context, s3Event *events.S3Event) {
	log.Printf("S3_BUCKET_FOLDER_IMAGES_COMPRESSED=%s S3_BUCKET_FOLDER_IMAGES_EXIF=%s S3_BUCKET_FOLDER_IMAGES_UPLOADED=%s",
		s3BucketFolderImagesCompressed,
		s3BucketFolderImagesExif,
		s3BucketFolderImagesUploaded)

	// Check if the S3 bucket folders are configured correctly.
	if s3BucketFolderImagesUploaded == s3BucketFolderImagesCompressed {
		log.Fatalf("S3_BUCKET_FOLDER_IMAGES_COMPRESSED == S3_BUCKET_FOLDER_IMAGES_UPLOADED")
	}

	// Create an AWS session and process the S3 event.
	awsSession := createAWSSession()
	processS3Event(awsSession, s3Event)
}

// main function is the entry point of the AWS Lambda application.
func main() {
	// Initialize S3 bucket folder variables from environment variables.
	s3BucketFolderImagesCompressed = getEnvironmentVariable("S3_BUCKET_FOLDER_IMAGES_COMPRESSED")
	s3BucketFolderImagesExif = getEnvironmentVariable("S3_BUCKET_FOLDER_IMAGES_EXIF")
	s3BucketFolderImagesUploaded = getEnvironmentVariable("S3_BUCKET_FOLDER_IMAGES_UPLOADED")

	// Start the AWS Lambda handler function.
	lambda.Start(handler)
}
