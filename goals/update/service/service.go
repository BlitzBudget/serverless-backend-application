package service

import (
	"fmt"
	"update-goal/service/repository"

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
		fmt.Printf("ParseToQueryParameter: Failed to marshal request to query parameter data %v \n", err)
		return err
	}

	av, err := repository.ParseToQueryParameter(request)
	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to marshal request to query parameter data %v \n", err)
		return err
	}

	err = repository.UpdateItem(av, svc, request)
	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to marshal request to query parameter data %v \n", err)
		return err
	}

	fmt.Printf("Successfully updated the item")
	return nil
}
