// Package main serves as the entry point for the AWS Rekognition image processing application.
package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Global variables to store S3 bucket folder names.
var (
	s3BucketFolderImagesCompressed        string
	s3BucketFolderRekognitionDetectFaces  string
	s3BucketFolderRekognitionDetectLabels string
	s3BucketFolderRekognitionDetectText   string
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
// It exits the program with an error if the variable is not set or empty.
func getEnvironmentVariable(key string) string {
	environmentValue := os.Getenv(key)
	if len(environmentValue) == 0 {
		log.Fatalf("%s is not set", key)
	}
	return environmentValue
}

// validateS3Folders checks that S3 bucket folder names are unique and not empty.
func validateS3Folders() {
	folders := []string{
		s3BucketFolderImagesCompressed,
		s3BucketFolderRekognitionDetectFaces,
		s3BucketFolderRekognitionDetectLabels,
		s3BucketFolderRekognitionDetectText,
	}

	// Create a map to store folder names and check for duplicates.
	folderMap := make(map[string]bool)

	for _, folder := range folders {
		if folder == "" {
			log.Fatal("S3 bucket folder names cannot be empty.")
		}
		if folderMap[folder] {
			log.Fatalf("Duplicate S3 bucket folder name found: %s", folder)
		}
		folderMap[folder] = true
	}
}

// handler is the AWS Lambda function that processes S3 events.
func handler(context context.Context, s3Event *events.S3Event) {
	log.Printf("S3_BUCKET_FOLDER_IMAGES_COMPRESSED=%s S3_BUCKET_FOLDER_REKOGNITION_DETECT_FACES=%s S3_BUCKET_FOLDER_REKOGNITION_DETECT_LABELS=%s S3_BUCKET_FOLDER_REKOGNITION_DETECT_TEXT=%s",
		s3BucketFolderImagesCompressed,
		s3BucketFolderRekognitionDetectFaces,
		s3BucketFolderRekognitionDetectLabels,
		s3BucketFolderRekognitionDetectText)

	// Create an AWS session and process the S3 event.
	awsSession := createAWSSession()
	processS3Event(awsSession, s3Event)
}

// main function is the entry point of the application.
func main() {
	// Initialize S3 bucket folder variables from environment variables.
	s3BucketFolderImagesCompressed = getEnvironmentVariable("S3_BUCKET_FOLDER_IMAGES_COMPRESSED")
	s3BucketFolderRekognitionDetectFaces = getEnvironmentVariable("S3_BUCKET_FOLDER_REKOGNITION_DETECT_FACES")
	s3BucketFolderRekognitionDetectLabels = getEnvironmentVariable("S3_BUCKET_FOLDER_REKOGNITION_DETECT_LABELS")
	s3BucketFolderRekognitionDetectText = getEnvironmentVariable("S3_BUCKET_FOLDER_REKOGNITION_DETECT_TEXT")

	// Validate S3 folder names.
	validateS3Folders()

	// Start the AWS Lambda handler function.
	lambda.Start(handler)
}
