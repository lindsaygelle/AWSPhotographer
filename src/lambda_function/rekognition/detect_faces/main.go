package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	_ = session.Must(session.NewSessionWithOptions(session.Options{}))
)

func handler(context context.Context, event interface{}) {
	switch event := event.(type) {
	case events.SNSEvent:
		handlerSNSEvent(event)
	}
}

func handlerSNSEvent(event events.SNSEvent) {
	for _, record := range event.Records {
		processSNSEventRecord(record.SNS)
	}
}

func processSNSEventRecord(record events.SNSEntity) {
	var s3Event events.S3Event
	err := json.Unmarshal([]byte(record.Message), &s3Event)
	if err != nil {
		fmt.Println(err)
		return
	}
	processS3Event(s3Event)
}

func processS3Event(event events.S3Event) {
	for _, record := range event.Records {
		processS3EventRecord(record.S3)
	}
}

func processS3EventRecord(record events.S3Entity) {
	processS3Object(record.Bucket, record.Object)
}

func processS3Object(s3Bucket events.S3Bucket, s3Object events.S3Object) {
	fmt.Println(s3Bucket.Name)
	fmt.Println(s3Object.ETag, s3Object.Key, s3Object.Size)
}

func main() {
	lambda.Start(handler)
}
