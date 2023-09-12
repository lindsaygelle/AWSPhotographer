package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

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
	processRekognition(session, s3BucketName, s3Object.Key)
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
