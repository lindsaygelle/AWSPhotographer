package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// getS3Object downloads a file from the provided S3 bucket using the AWS S3 Downloader.
// It saves the file with the given fileName and returns any error encountered.
func getS3Object(s3DownloadManager *s3manager.Downloader, s3BucketName string, s3ObjectKey string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	getObjectInput := s3.GetObjectInput{
		Bucket: &s3BucketName,
		Key:    &s3ObjectKey}
	_, err = s3DownloadManager.Download(file, &getObjectInput)
	if err != nil {
		return err
	}
	return nil
}

// putS3ObjectExifMetadata uploads Exif metadata to the specified S3 bucket.
// It serializes the ExifMetadata to JSON and stores it in the S3 object.
// Returns the S3 PutObjectOutput and any error encountered.
func putS3ObjectExifMetadata(s3Client *s3.S3, s3BucketName string, s3ObjectKey string, exifMetadata *ExifMetadata) (*s3.PutObjectOutput, error) {
	b, err := json.Marshal(exifMetadata)
	if err != nil {
		return nil, err
	}
	s3PutObjectInput := s3.PutObjectInput{
		Bucket:        &s3BucketName,
		Body:          aws.ReadSeekCloser(bytes.NewReader(b)),
		ContentLength: aws.Int64(int64(len(b))),
		ContentType:   aws.String("json"),
		Key:           &s3ObjectKey}
	s3PutObjectOutput, err := s3Client.PutObject(&s3PutObjectInput)
	if err != nil {
		return nil, err
	}
	return s3PutObjectOutput, err
}

// putS3ObjectImageJpg uploads an image in JPEG format to the specified S3 bucket.
// Returns the S3 PutObjectOutput and any error encountered.
func putS3ObjectImageJpg(s3UploadManager *s3manager.Uploader, s3BucketName string, s3ObjectKey string, sourceImage image.Image) (*s3manager.UploadOutput, error) {
	var buffer bytes.Buffer
	err := jpeg.Encode(&buffer, sourceImage, nil)
	if err != nil {
		return nil, err
	}
	s3ManagerUploadInput := s3manager.UploadInput{
		Body:        bytes.NewReader(buffer.Bytes()),
		Bucket:      &s3BucketName,
		ContentType: aws.String("image/jpeg"),
		Key:         &s3ObjectKey}
	s3ManagerUploadOutput, err := s3UploadManager.Upload(&s3ManagerUploadInput)
	if err != nil {
		return nil, err
	}
	return s3ManagerUploadOutput, err
}
