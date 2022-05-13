package service

import (
	"fmt"
	"get-budget/service/models"
	"get-budget/service/repository"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func QueryItems(body *string) []*models.ResponseItem {
	// snippet-start:[dynamodb.go.create_item.session]
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	// snippet-end:[dynamodb.go.create_item.session]

	av, err := repository.AttributeBuilder(body)
	if err != nil {
		log.Fatalf("Got error marshalling new item: %v", err)
		return nil
	}

	var queryOutput *dynamodb.QueryOutput
	queryOutput, err = repository.QueryItem(av, svc)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %v", err)
		return nil
	}

	var responseItems []*models.ResponseItem
	responseItems, err = repository.ParseResponse(queryOutput)
	if err != nil {
		log.Fatalf("Got error parsing Response Item: %v", err)
		return nil
	}

	fmt.Printf("Successfully retrieved %v items with the consumed capacity of %v'", queryOutput.Count, queryOutput.ConsumedCapacity)
	return responseItems
}
