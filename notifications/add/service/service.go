package service

import (
	"add-notification/service/repository"
	"fmt"

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

	av, err := repository.AttributeBuilder(body)
	if err != nil {
		fmt.Printf("SaveRequest: Got error marshalling new item: %v", err)
		return err
	}

	err = repository.CreateItem(av, svc)
	if err != nil {
		fmt.Printf("SaveRequest: Got error calling PutItem: %v", err)
		return err
	}

	fmt.Println("Successfully added the item!")
	return nil
}
