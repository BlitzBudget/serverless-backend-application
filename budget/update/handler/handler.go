package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"patch-budget/service"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jsonReq, _ := json.Marshal(request)
	fmt.Printf("Processing request data for request %v.\n", string(jsonReq))

	service.SaveRequest(&request.Body)
	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,PATCH",
		"Access-Control-Allow-Credentials": "true",
	}
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200, Headers: header}, nil
}
