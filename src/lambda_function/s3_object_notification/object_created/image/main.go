package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(context context.Context, event *events.S3Event) {
	for index, record := range event.Records {
		fmt.Println(index, record)
	}
}

func main() {
	lambda.Start(handler)
}
