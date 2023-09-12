package main

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getS3PutObjectInput(s3BucketName string, s3ObjectKey string, s3ObjectBody interface{}) (*s3.PutObjectInput, error) {
	b, err := json.Marshal(s3ObjectBody)
	if err != nil {
		return nil, err
	}
	s3PutObjectInput := s3.PutObjectInput{
		Bucket:        &s3BucketName,
		Body:          aws.ReadSeekCloser(bytes.NewReader(b)),
		ContentLength: aws.Int64(int64(len(b))),
		ContentType:   aws.String("json"),
		Key:           &s3ObjectKey}
	return &s3PutObjectInput, nil
}

func putS3Object(s3Client *s3.S3, s3PutObjectInput *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	s3PutObjectOutput, err := s3Client.PutObject(s3PutObjectInput)
	if err != nil {
		return nil, err
	}
	return s3PutObjectOutput, err
}
