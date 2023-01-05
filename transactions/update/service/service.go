package service

import (
	"fmt"
	"patch-transactions/service/repository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func SaveRequest(body *string) error {
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

	request, err := repository.AttributeBuilder(body)
	if err != nil {
		fmt.Printf("SaveRequest: There was an error unmarshalling the bytes to struct: %v \n", err.Error())
		return nil
	}

	av, err := repository.ParseToQueryParameter(request)
	if err != nil {
		fmt.Printf("SaveRequest: Parsing body to Query Parameter: %v \n", err.Error())
		return nil
	}

	err = repository.UpdateItem(av, svc, request)
	if err != nil {
		fmt.Printf("SaveRequest: There is an error Updating Item: %v \n", err.Error())
		return nil
	}

	fmt.Printf("Successfully updated the item'")
	return nil
}
