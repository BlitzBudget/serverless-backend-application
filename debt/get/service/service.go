package service

import (
	"fmt"
	"get-debt/service/models"
	"get-debt/service/repository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func QueryItems(body *string) ([]*models.ResponseItem, error) {
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
		fmt.Printf("Got error marshalling new item: %v \n", err)
		return nil, err
	}

	var queryOutput *dynamodb.QueryOutput
	queryOutput, err = repository.QueryItem(av, svc)
	if err != nil {
		fmt.Printf("Got error calling PutItem: %v \n", err)
		return nil, err
	}

	var responseItems []*models.ResponseItem
	responseItems, err = repository.ParseResponse(queryOutput)
	if err != nil {
		fmt.Printf("Got error parsing Response Item: %v \n", err)
		return nil, err
	}

	fmt.Printf("Successfully retrieved %v items with the consumed capacity of %v'", *queryOutput.Count, queryOutput.ConsumedCapacity)
	return responseItems, nil
}
