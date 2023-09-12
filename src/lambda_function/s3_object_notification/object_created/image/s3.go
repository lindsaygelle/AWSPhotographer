package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3DownloadManager interface {
	Download(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error)
}

type S3UploadManager interface {
	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

// getS3ObjectExifMetadata downloads a file from an S3 bucket and extracts its exif metadata.
func getS3ObjectExifMetadata(session *session.Session, bucket *events.S3Bucket, object *events.S3Object) {
	if !strings.HasPrefix(object.Key, s3BucketFolderImagesUploaded) {
		log.Fatalf("%s does not begin with %s", object.Key, s3BucketFolderImagesUploaded)
	}
	filePath, err := getS3Object(s3manager.NewDownloader(session), fileDirectory, path.Base(object.Key), bucket.Name, object.Key)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalln(err)
	}
	exifMetadata := getExif(file)
	b, err := json.Marshal(&exifMetadata)
	if err != nil {
		log.Fatalln(err)
	}
	fileName := strings.Split(path.Base(*filePath), ".")[0]
	s3PutObjectOutput, err := putS3Object(s3.New(session), bucket.Name, fmt.Sprintf("%s/%s.json", s3BucketFolderImagesExif, fileName), &b)
	if err != nil {
		log.Fatalln(err)
	}
	processS3PutObjectOutput(s3PutObjectOutput)
}

// getS3Object downloads a file from the provided S3 bucket.
func getS3Object(s3DownloadManager S3DownloadManager, fileDirectory string, fileName string, bucketName string, bucketObjectKey string) (*string, error) {
	file, err := os.Create(fmt.Sprintf("%s/%s", fileDirectory, fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	getObjectInput := s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &bucketObjectKey}
	_, err = s3DownloadManager.Download(file, &getObjectInput)
	if err != nil {
		return nil, err
	}
	filePath := file.Name()
	return &filePath, nil
}

// putS3Object uploads a file (or at least its contents) to the provided S3 bucket.
func putS3Object(s3UploadManager S3UploadManager, bucketName string, bucketObjectKey string, b *[]byte) (*s3.PutObjectOutput, error) {
	s3PutObjectInput := s3.PutObjectInput{
		Bucket: &bucketName,
		Body:   aws.ReadSeekCloser(bytes.NewReader(*b)),
		Key:    &bucketObjectKey}
	s3PutObjectOutput, err := s3UploadManager.PutObject(&s3PutObjectInput)
	if err != nil {
		return nil, err
	}
	return s3PutObjectOutput, err
}

// processS3PutObjectOutput process an s3.PutObjectOutput.
func processS3PutObjectOutput(s3PutObjectOutput *s3.PutObjectOutput) {
	log.Printf("S3PutObjectOutput: BucketKeyEnabled=%v ChecksumCRC32=%v ChecksumCRC32C=%v ChecksumSHA1=%v ChecksumSHA256=%v ETag=%v Expiration=%v RequestCharged=%v SSECustomerAlgorithm=%v SSECustomerKeyMD5=%v SSEKMSEncryptionContext=%v SSEKMSKeyId=%v ServerSideEncryption=%v VersionId=%v",
		s3PutObjectOutput.BucketKeyEnabled,
		s3PutObjectOutput.ChecksumCRC32,
		s3PutObjectOutput.ChecksumCRC32C,
		s3PutObjectOutput.ChecksumSHA1,
		s3PutObjectOutput.ChecksumSHA256,
		s3PutObjectOutput.ETag,
		s3PutObjectOutput.Expiration,
		s3PutObjectOutput.RequestCharged,
		s3PutObjectOutput.SSECustomerAlgorithm,
		s3PutObjectOutput.SSECustomerKeyMD5,
		s3PutObjectOutput.SSEKMSEncryptionContext,
		s3PutObjectOutput.SSEKMSKeyId,
		s3PutObjectOutput.ServerSideEncryption,
		s3PutObjectOutput.VersionId)
}

// processS3Bucket processes an events.S3Bucket.
func processS3Bucket(bucket *events.S3Bucket) {
	log.Printf("S3Bucket: Arn=%s Name=%s OwnerIdentity.PrincipalID=%s",
		bucket.Arn,
		bucket.Name,
		bucket.OwnerIdentity.PrincipalID)
}

// processS3Event process an events.S3Event.
func processS3Event(session *session.Session, event *events.S3Event) {
	log.Printf("S3Event: Records=%d", len(event.Records))
	for _, eventRecord := range event.Records {
		log.Printf("S3EventRecord: Index=%d", len(event.Records))
		processS3EventRecord(session, &eventRecord)
	}
}

// processS3EventRecord processes an events.S3EventRecord.
func processS3EventRecord(session *session.Session, eventRecord *events.S3EventRecord) {
	log.Printf("S3EventRecord: AWSRegion=%s EventTime=%s EventName=%s EventSource=%s EventVersion=%s",
		eventRecord.AWSRegion,
		eventRecord.EventTime,
		eventRecord.EventName,
		eventRecord.EventSource,
		eventRecord.EventVersion)
	processS3Entity(session, &eventRecord.S3)
}

// processS3Entity process an events.S3Entity.
func processS3Entity(session *session.Session, entity *events.S3Entity) {
	log.Printf("S3Entity: ConfigurationID=%s SchemaVersion=%s",
		entity.ConfigurationID,
		entity.SchemaVersion)
	processS3Bucket(&entity.Bucket)
	processS3Object(&entity.Object)
	getS3ObjectExifMetadata(session, &entity.Bucket, &entity.Object)
}

// processS3Object processes an events.S3Object.
func processS3Object(object *events.S3Object) {
	log.Printf("S3Object: ETag=%s Key=%s Sequencer=%s Size=%d URLDecodeKey=%s VersionID=%s",
		object.ETag,
		object.Key,
		object.Sequencer,
		object.Size,
		object.URLDecodedKey,
		object.VersionID)
}
