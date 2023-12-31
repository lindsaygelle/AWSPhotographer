package main

import (
	"fmt"
	"log"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Contains functions for processing AWS S3 events and managing S3 objects.

// createS3ObjectKey generates an S3 object key based on the provided parameters.
func createS3ObjectKey(s3BucketName string, fileName string, fileTime time.Time) string {
	return fmt.Sprintf("%s/%s/%s", s3BucketName, fileTime.Format("2006/01/02"), fileName)
}

// processS3PutObjectOutput processes the output of an S3 PutObject operation.
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

// processS3ManagerUploadOutput processes the output of an S3 Upload operation.
func processS3ManagerUploadOutput(s3ManagerUploadOutput *s3manager.UploadOutput) {
	log.Printf("S3managerUploadOutput: ETag=%v Location=%v UploadID=%v VersionID=%v",
		s3ManagerUploadOutput.ETag,
		s3ManagerUploadOutput.Location,
		s3ManagerUploadOutput.UploadID,
		s3ManagerUploadOutput.VersionID)
}

// processS3Bucket processes an AWS S3 bucket event.
func processS3Bucket(session *session.Session, s3Bucket *events.S3Bucket) {
	log.Printf("S3Bucket: Arn=%s Name=%s OwnerIdentity.PrincipalID=%s",
		s3Bucket.Arn,
		s3Bucket.Name,
		s3Bucket.OwnerIdentity.PrincipalID)
}

// processS3Event processes an AWS S3 event.
func processS3Event(session *session.Session, s3Event *events.S3Event) {
	log.Printf("S3Event: Records=%d", len(s3Event.Records))
	for index, s3EventRecord := range s3Event.Records {
		log.Printf("S3EventRecord: Index=%d", index)
		processS3EventRecord(session, &s3EventRecord)
	}
}

// processS3EventRecord processes an AWS S3 event record.
func processS3EventRecord(session *session.Session, s3EventRecord *events.S3EventRecord) {
	log.Printf("S3EventRecord: AWSRegion=%s EventTime=%s EventName=%s EventSource=%s EventVersion=%s",
		s3EventRecord.AWSRegion,
		s3EventRecord.EventTime,
		s3EventRecord.EventName,
		s3EventRecord.EventSource,
		s3EventRecord.EventVersion)
	processS3Entity(session, &s3EventRecord.S3)
}

// processS3Entity processes an AWS S3 entity.
func processS3Entity(session *session.Session, s3Entity *events.S3Entity) {
	log.Printf("S3Entity: ConfigurationID=%s SchemaVersion=%s",
		s3Entity.ConfigurationID,
		s3Entity.SchemaVersion)
	processS3Bucket(session, &s3Entity.Bucket)
	processS3Object(session, s3Entity.Bucket.Name, &s3Entity.Object)
}

// processS3Object processes an AWS S3 object event.
func processS3Object(session *session.Session, s3BucketName string, s3Object *events.S3Object) {
	log.Printf("S3Object: Bucket=%s ETag=%s Key=%s Sequencer=%s Size=%d URLDecodeKey=%s VersionID=%s",
		s3BucketName,
		s3Object.ETag,
		s3Object.Key,
		s3Object.Sequencer,
		s3Object.Size,
		s3Object.URLDecodedKey,
		s3Object.VersionID)

	// Construct the local file name for the S3 object.
	fileName := fmt.Sprintf("%s/%s", fileDirectory, path.Base(s3Object.Key))

	// Download the S3 object to the local file.
	err := getS3Object(s3manager.NewDownloader(session), s3BucketName, s3Object.Key, fileName)
	if err != nil {
		log.Fatalf("S3Object: Bucket=%s Key=%s FileName=%s Error=%s", s3BucketName, s3Object.Key, fileName, err)
	}

	// Create an S3 client.
	s3Client := s3.New(session)

	// Process Exif metadata for the S3 object.
	processS3ObjectExifMetadata(session, s3Client, s3BucketName, fileName)
}

// processS3ObjectExifMetadata processes Exif metadata for an AWS S3 object event.
func processS3ObjectExifMetadata(session *session.Session, s3Client *s3.S3, s3BucketName string, fileName string) {
	log.Printf("ExifMetadata: BucketName=%s FileName=%s", s3BucketName, fileName)

	// Open and extract Exif metadata from the image file.
	exifMetadata, err := openExif(fileName)
	if err != nil {
		log.Fatalf("ExifMetadata: Error=%s", err)
	}
	log.Println("ExifMetata: Successfully created ExifMetadata")

	// Create the S3 object key for the Exif metadata.
	s3ObjectKey := createS3ObjectKey(s3BucketFolderImagesExif, fmt.Sprintf("%s.JSON", strings.Split(path.Base(fileName), ".")[0]), *exifMetadata.DateTime)
	log.Printf("ExifMetadata: Bucket=%s Key=%s", s3BucketName, s3ObjectKey)

	// Upload the Exif metadata to S3.
	s3PutObjectOutput, err := putS3ObjectExifMetadata(s3Client, s3BucketName, s3ObjectKey, exifMetadata)
	if err != nil {
		log.Fatalf("ExifMetadata: Error=%s", err)
	}
	log.Println("ExifMetata: Successfully uploaded ExifMetadata")

	// Process the S3 PutObject output.
	processS3PutObjectOutput(s3PutObjectOutput)

	// Process the image itself.
	processS3ObjectImage(session, s3Client, s3BucketName, fileName, *exifMetadata.DateTime)
}

// processS3ObjectImage processes an image for an AWS S3 object event.
func processS3ObjectImage(session *session.Session, s3Client *s3.S3, s3BucketName string, fileName string, fileTime time.Time) {
	log.Printf("CompressImage: BucketName=%s FileName=%s", s3BucketName, fileName)

	// Open and prepare the image for compression.
	image, err := openImage(fileName)
	if err != nil {
		log.Fatalf("CompressImage: Error=%s", err)
	}
	log.Println("CompressImage: Successfully created image")

	// Ensure that the compressed folder configuration is correct.
	if strings.Contains(s3BucketFolderImagesCompressed, "upload") {
		log.Panicf("Compressed folder is incorrect! %s", s3BucketFolderImagesCompressed)
	}

	// Create the S3 object key for the compressed image.
	s3ObjectKey := createS3ObjectKey(s3BucketFolderImagesCompressed, path.Base(fileName), fileTime)
	log.Printf("CompressImage: Bucket=%s Key=%s", s3BucketName, s3ObjectKey)

	// Upload the compressed image to S3.
	s3ManagerUploadOutput, err := putS3ObjectImageJpg(s3manager.NewUploader(session), s3BucketName, s3ObjectKey, image)
	if err != nil {
		log.Fatalf("Image: Error=%s", err)
	}
	log.Println("CompressImage: Successfully uploaded image")

	// Process the S3 Upload output.
	processS3ManagerUploadOutput(s3ManagerUploadOutput)
}
