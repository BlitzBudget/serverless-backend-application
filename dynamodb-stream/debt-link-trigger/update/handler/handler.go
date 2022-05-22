package handler

import (
	"add-debt-link/service"
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.DynamoDBEvent) error {
	fmt.Printf("Processing request data for request %v.\n", request.Records)
	fmt.Printf("Record Length = %v.\n", len(request.Records))
	fmt.Println("Records:")
	for key, value := range request.Records {
		fmt.Printf("    %v: %v\n", key, value.EventName)
	}

	service.ProcessRecords(&request.Records)

	return nil
}
