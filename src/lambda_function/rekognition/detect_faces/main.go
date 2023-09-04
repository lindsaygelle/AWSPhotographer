package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		Body:            "\"Hello from Lambda!\"",
		IsBase64Encoded: false,
		StatusCode:      200,
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
