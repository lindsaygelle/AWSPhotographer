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

// processRekognition processes an image object in S3 against various AWS Rekognition services.
func processRekognition(session *session.Session, s3BucketName string, s3ObjectKey string) {
	rekognitionClient := rekognition.New(session)
	s3Client := s3.New(session)
	processRekognitionDetectFaces(rekognitionClient, s3Client, s3BucketName, s3ObjectKey)
	processRekognitionDetectLabels(rekognitionClient, s3Client, s3BucketName, s3ObjectKey)
	processRekognitionDetectText(rekognitionClient, s3Client, s3BucketName, s3ObjectKey)
}

// processRekognitionDetectFaces processes an S3 object image with AWS Rekognition Detect Faces.
func processRekognitionDetectFaces(rekognitionClient *rekognition.Rekognition, s3Client *s3.S3, s3BucketName string, s3ObjectKey string) {
	// Create a Rekognition S3Object representation.
	rekognitionS3Object := rekognition.S3Object{
		Bucket: &s3BucketName,
		Name:   &s3ObjectKey}
	processRekognitionS3Object(&rekognitionS3Object)

	// Create a Rekognition Image representation.
	rekognitionImage := rekognition.Image{
		S3Object: &rekognitionS3Object,
	}
	processRekognitionImage(&rekognitionImage)

	// Create a Rekognition DetectFacesInput and process it.
	rekognitionDetectFacesInput := rekognition.DetectFacesInput{
		Attributes: aws.StringSlice([]string{"ALL"}),
		Image:      &rekognitionImage,
	}
	processRekognitionDetectFacesInput(&rekognitionDetectFacesInput)

	// Perform face detection using Rekognition.
	rekognitionDetectFacesOutput, err := rekognitionClient.DetectFaces(&rekognitionDetectFacesInput)
	if err != nil {
		log.Fatalf("RekognitionDetectFacesOutput: Error=%s", err)
	}

	// Process the output and store it in S3.
	processRekognitionDetectFacesOutput(s3Client, rekognitionDetectFacesOutput, s3BucketName, s3ObjectKey)
}

// processRekognitionDetectFacesInput processes a rekognition.DetectFacesInput.
func processRekognitionDetectFacesInput(rekognitionDetectFacesInput *rekognition.DetectFacesInput) {
	log.Printf("RekognitionDetectFacesInput: Attributes=%v Bytes=%v",
		rekognitionDetectFacesInput.Attributes,
		rekognitionDetectFacesInput.Image.Bytes)
}

// processRekognitionDetectFacesOutput processes a rekognition.DetectFacesOutput.
func processRekognitionDetectFacesOutput(s3Client *s3.S3, rekognitionDetectFacesOutput *rekognition.DetectFacesOutput, s3BucketName string, s3ObjectKey string) {
	log.Printf("RekognitionDetectFacesOutput: OrientationCorrection=%v",
		rekognitionDetectFacesOutput.OrientationCorrection)

	// Modify the S3 object key for storage.
	s3ObjectKey = fmt.Sprintf("%s/%s.JSON", s3BucketFolderRekognitionDetectFaces, strings.Split(path.Base(s3ObjectKey), ".")[0])

	// Prepare S3 PutObjectInput and perform the S3 object update.
	s3PutObjectInput, err := getS3PutObjectInput(s3BucketName, s3ObjectKey, rekognitionDetectFacesOutput)
	if err != nil {
		log.Fatalf("S3PutObjectInput: Error=%s", err)
	}

	// Todo: Implement process S3PutObjectInput
	s3PutObjectOutput, err := putS3Object(s3Client, s3PutObjectInput)
	if err != nil {
		log.Fatalf("S3PutObjectOutput: Error=%s", err)
	}
	processS3PutObjectOutput(s3PutObjectOutput)
}

// processRekognitionDetectLabels processes an S3 object image with AWS Rekognition Detect Labels.
func processRekognitionDetectLabels(rekognitionClient *rekognition.Rekognition, s3Client *s3.S3, s3BucketName string, s3ObjectKey string) {
	// Create a Rekognition S3Object representation.
	rekognitionS3Object := rekognition.S3Object{
		Bucket: &s3BucketName,
		Name:   &s3ObjectKey}
	processRekognitionS3Object(&rekognitionS3Object)

	// Create a Rekognition Image representation.
	rekognitionImage := rekognition.Image{
		S3Object: &rekognitionS3Object,
	}
	processRekognitionImage(&rekognitionImage)

	// Create a Rekognition DetectLabelsInput and process it.
	rekognitionDetectLabelsInput := rekognition.DetectLabelsInput{
		Features: aws.StringSlice([]string{"GENERAL_LABELS", "IMAGE_PROPERTIES"}),
		Image:    &rekognitionImage,
	}
	processRekognitionDetectLabelsInput(&rekognitionDetectLabelsInput)

	// Perform label detection using Rekognition.
	rekognitionDetectLabelsOutput, err := rekognitionClient.DetectLabels(&rekognitionDetectLabelsInput)
	if err != nil {
		log.Fatalf("RekognitionDetectLabelsOutput: Error=%s", err)
	}

	// Process the output and store it in S3.
	processRekognitionDetectLabelsOutput(s3Client, rekognitionDetectLabelsOutput, s3BucketName, s3ObjectKey)
}

// processRekognitionDetectLabelsInput processes a rekognition.DetectLabelsInput.
func processRekognitionDetectLabelsInput(rekognitionDetectLabelsInput *rekognition.DetectLabelsInput) {
	log.Printf("RekognitionDetectLabelsInput: Features=%v Bytes=%v MaxLabels=%v",
		rekognitionDetectLabelsInput.Features,
		rekognitionDetectLabelsInput.Image.Bytes,
		rekognitionDetectLabelsInput.MaxLabels)
}

// processRekognitionDetectLabelsOutput processes a rekognition.DetectLabelsOutput.
func processRekognitionDetectLabelsOutput(s3Client *s3.S3, rekognitionDetectLabelsOutput *rekognition.DetectLabelsOutput, s3BucketName string, s3ObjectKey string) {
	log.Printf("RekognitionDetectLabelsOutput: OrientationCorrection=%v",
		rekognitionDetectLabelsOutput.OrientationCorrection)

	// Modify the S3 object key for storage.
	s3ObjectKey = fmt.Sprintf("%s/%s.JSON", s3BucketFolderRekognitionDetectLabels, strings.Split(path.Base(s3ObjectKey), ".")[0])

	// Prepare S3 PutObjectInput and perform the S3 object update.
	s3PutObjectInput, err := getS3PutObjectInput(s3BucketName, s3ObjectKey, rekognitionDetectLabelsOutput)
	if err != nil {
		log.Fatalf("S3PutObjectInput: Error=%s", err)
	}

	// Perform S3 PutObject operation to store the processed data.
	s3PutObjectOutput, err := putS3Object(s3Client, s3PutObjectInput)
	if err != nil {
		log.Fatalf("S3PutObjectOutput: Error=%s", err)
	}
	processS3PutObjectOutput(s3PutObjectOutput)
}

// processRekognitionDetectText processes an S3 object image with AWS Rekognition Detect Faces.
func processRekognitionDetectText(rekognitionClient *rekognition.Rekognition, s3Client *s3.S3, s3BucketName string, s3ObjectKey string) {
	// Create a Rekognition S3Object representation.
	rekognitionS3Object := rekognition.S3Object{
		Bucket: &s3BucketName,
		Name:   &s3ObjectKey}
	processRekognitionS3Object(&rekognitionS3Object)

	// Create a Rekognition Image representation.
	rekognitionImage := rekognition.Image{
		S3Object: &rekognitionS3Object,
	}
	processRekognitionImage(&rekognitionImage)

	// Create a Rekognition DetectTextInput and process it.
	rekognitionDetectTextInput := rekognition.DetectTextInput{
		Image: &rekognitionImage,
	}
	processRekognitionDetectTextInput(&rekognitionDetectTextInput)

	// Perform face detection using Rekognition.
	rekognitionDetectTextOutput, err := rekognitionClient.DetectText(&rekognitionDetectTextInput)
	if err != nil {
		log.Fatalf("RekognitionDetectTextOutput: Error=%s", err)
	}

	// Process the output and store it in S3.
	processRekognitionDetectTextOutput(s3Client, rekognitionDetectTextOutput, s3BucketName, s3ObjectKey)
}

// processRekognitionDetectTextInput processes a rekognition.DetectTextInput.
func processRekognitionDetectTextInput(rekognitionDetectTextInput *rekognition.DetectTextInput) {
	log.Printf("RekognitionDetectTextInput: Filters=%v",
		rekognitionDetectTextInput.Filters)
}

// processRekognitionDetectTextOutput processes a rekognition.DetectTextOutput.
func processRekognitionDetectTextOutput(s3Client *s3.S3, rekognitionDetectTextOutput *rekognition.DetectTextOutput, s3BucketName string, s3ObjectKey string) {
	log.Printf("RekognitionDetectTextOutput: TextModelVersion =%v",
		rekognitionDetectTextOutput.TextModelVersion)

	// Modify the S3 object key for storage.
	s3ObjectKey = fmt.Sprintf("%s/%s.JSON", s3BucketFolderRekognitionDetectText, strings.Split(path.Base(s3ObjectKey), ".")[0])

	// Prepare S3 PutObjectInput and perform the S3 object update.
	s3PutObjectInput, err := getS3PutObjectInput(s3BucketName, s3ObjectKey, rekognitionDetectTextOutput)
	if err != nil {
		log.Fatalf("S3PutObjectInput: Error=%s", err)
	}

	// Todo: Implement process S3PutObjectInput
	s3PutObjectOutput, err := putS3Object(s3Client, s3PutObjectInput)
	if err != nil {
		log.Fatalf("S3PutObjectOutput: Error=%s", err)
	}
	processS3PutObjectOutput(s3PutObjectOutput)
}

// processRekognitionS3Object processes a rekognition.S3Object.
func processRekognitionS3Object(rekognitionS3Object *rekognition.S3Object) {
	log.Printf("RekognitionS3Object: Bucket=%s Key=%s",
		*rekognitionS3Object.Bucket,
		*rekognitionS3Object.Name)
}

// processRekognitionImage processes a rekognition.Image.
func processRekognitionImage(rekognitionImage *rekognition.Image) {
	log.Printf("RekognitionImage: S3Object.Bucket=%s S3Object.Name=%s",
		*rekognitionImage.S3Object.Bucket,
		*rekognitionImage.S3Object.Name)
}
