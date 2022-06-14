package handler

import (
	"add-goal-link/service"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.DynamoDBEvent) error {
	b, err := json.Marshal(request.Records)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	fmt.Printf("Processing request data for request %v.\n", request.Records)
	fmt.Printf("Record Length = %v.\n", len(request.Records))
	fmt.Println("Records:")
	for key, value := range request.Records {
		fmt.Printf("    %v: %v\n", key, value.EventName)
	}

	service.ProcessRecords(&request.Records)

	return nil
}
