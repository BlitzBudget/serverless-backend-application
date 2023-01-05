package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"get-tag/service"
	"get-tag/service/models"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %v.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %v.\n", request.Body)
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %v: %v\n", key, value)
	}

	data, err := service.QueryItems(&request.Body)
	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	body, errParsing := json.Marshal(data)
	if errParsing != nil {
		fmt.Printf("HandleRequest:: Failed to unmarshal Response data %v", errParsing)
		errorMessage = errParsing.Error()
	}

	if errorMessage != "" {
		errorRespose := models.ErrorHttpResponse{
			Message: &errorMessage,
		}
		errorAsBytes, _ := json.Marshal(errorRespose)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: 500, Headers: header}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200, Headers: header}, nil
}
