package handler

import (
	"add-transactions/service"
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %v.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %v.\n", request.Body)
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %v: %v\n", key, value)
	}

	service.SaveRequest(&request.Body)
	header := map[string]string{
		"Access-Control-Allow-Origin":      "http://constellation-technology.com",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200, Headers: header}, nil
}
