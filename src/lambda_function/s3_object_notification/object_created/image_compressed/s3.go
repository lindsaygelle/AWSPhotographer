package main

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Contains utility functions for interacting with Amazon S3, specifically for putting objects.

// getS3PutObjectInput creates an S3 PutObjectInput based on the provided parameters.
// It serializes the s3ObjectBody to JSON and prepares the necessary input for object storage.
func getS3PutObjectInput(s3BucketName string, s3ObjectKey string, s3ObjectBody interface{}) (*s3.PutObjectInput, error) {
	// Serialize the s3ObjectBody to JSON.
	b, err := json.Marshal(s3ObjectBody)
	if err != nil {
		return nil, err
	}

	// Create an S3 PutObjectInput with the required fields.
	s3PutObjectInput := s3.PutObjectInput{
		Bucket:        &s3BucketName,
		Body:          aws.ReadSeekCloser(bytes.NewReader(b)),
		ContentLength: aws.Int64(int64(len(b))),
		ContentType:   aws.String("application/json"),
		Key:           &s3ObjectKey,
	}

	return &s3PutObjectInput, nil
}

// putS3Object puts an object into an Amazon S3 bucket using the provided S3 client and input.
// It returns the result of the PutObject operation or an error if the operation fails.
func putS3Object(s3Client *s3.S3, s3PutObjectInput *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	// Perform the S3 PutObject operation.
	s3PutObjectOutput, err := s3Client.PutObject(s3PutObjectInput)
	if err != nil {
		return nil, err
	}

	return s3PutObjectOutput, nil
}
