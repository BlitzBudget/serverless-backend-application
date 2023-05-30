package handler

import (
	"add-goal-link/service"
	"context"
	"encoding/json"
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
	b, err := json.Marshal(request.Records)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("JSON String is %v. \n", string(b))

	service.ProcessRecords(&request.Records)

	return nil
}
