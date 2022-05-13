package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"get-goal/service"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %v.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %v.\n", request.Body)
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %v: %v\n", key, value)
	}

	data := service.QueryItems(&request.Body)
	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}
	body, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Response data %v", err))
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200, Headers: header}, nil
}
