package main

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
)

func processRekognition(session *session.Session, s3BucketName string, s3ObjectKey string) {
	rekognitionClient := rekognition.New(session)
	s3Client := s3.New(session)
	processRekognitionDetectFaces(rekognitionClient, s3Client, s3BucketName, s3ObjectKey)
}

func processRekognitionDetectFaces(rekognitionClient *rekognition.Rekognition, s3Client *s3.S3, s3BucketName string, s3ObjectKey string) {
	rekognitionS3Object := rekognition.S3Object{
		Bucket: &s3BucketName,
		Name:   &s3ObjectKey}
	processRekognitionS3Object(&rekognitionS3Object)
	rekognitionImage := rekognition.Image{
		S3Object: &rekognitionS3Object,
	}
	processRekognitionImage(&rekognitionImage)
	rekognitionDetectFacesInput := rekognition.DetectFacesInput{
		Attributes: aws.StringSlice([]string{"ALL"}),
		Image:      &rekognitionImage}
	processRekognitionDetectFacesInput(&rekognitionDetectFacesInput)
	rekognitionDetectFacesOutput, err := rekognitionClient.DetectFaces(&rekognitionDetectFacesInput)
	if err != nil {
		log.Fatalf("RekognitionDetectFacesOuput: Error=%s", err)
	}
	processRekognitionDetectFacesOutput(s3Client, rekognitionDetectFacesOutput, s3BucketName, s3ObjectKey)
}

func processRekognitionDetectFacesInput(rekognitionDetectFacesInput *rekognition.DetectFacesInput) {
	log.Printf("RekognitionDetectFacesInput: Attributes=%v Bytes=%v",
		rekognitionDetectFacesInput.Attributes,
		rekognitionDetectFacesInput.Image.Bytes)
}

func processRekognitionDetectFacesOutput(s3Client *s3.S3, rekognitionDetectFacesOutput *rekognition.DetectFacesOutput, s3BucketName string, s3ObjectKey string) {
	log.Printf("RekognitionDetectFacesOutput: OrientationCorrection=%v",
		rekognitionDetectFacesOutput.OrientationCorrection)
	s3ObjectKey = fmt.Sprintf("%s/%s.JSON", s3BucketFolderRekognitionDetectFaces, strings.Split(path.Base(s3ObjectKey), ".")[0])
	s3PutObjectInput, err := getS3PutObjectInput(s3BucketName, s3ObjectKey, rekognitionDetectFacesOutput)
	if err != nil {
		log.Fatalf("S3PutObjectInput: Error=%s", err)
	}
	// Todo: Implement process S3PutObjectInput
	s3PutObjectOuput, err := putS3Object(s3Client, s3PutObjectInput)
	if err != nil {
		log.Fatalf("s3PutObjectOuput: Error=%s", err)
	}
	processS3PutObjectOutput(s3PutObjectOuput)
}

func processRekognitionS3Object(rekognitionS3Object *rekognition.S3Object) {
	log.Printf("RekognitionS3Object: Bucket=%s Key=%s",
		*rekognitionS3Object.Bucket,
		*rekognitionS3Object.Name)
}

func processRekognitionImage(rekognitionImage *rekognition.Image) {
	log.Printf("RekognitionImage: S3Object.Bucket=%s S3Object.Name=%s",
		*rekognitionImage.S3Object.Bucket,
		*rekognitionImage.S3Object.Name)
}
