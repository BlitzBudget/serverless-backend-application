package handler

import (
	"add-goals/service"
	"add-goals/service/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %v.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %v.\n", request.Body)
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf(" %v: %v\n", key, value)
	}

	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,PUT",
		"Access-Control-Allow-Credentials": "true",
	}

	err := service.SaveRequest(&request.Body)
	if err != nil {
		errorMessage := err.Error()
		errorRespose := models.ErrorHttpResponse{
			Message: &errorMessage,
		}
		errorAsBytes, _ := json.Marshal(errorRespose)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: 500, Headers: header}, nil
	}
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200, Headers: header}, nil
}
