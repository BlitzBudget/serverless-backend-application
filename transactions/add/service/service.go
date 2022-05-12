package service

import (
	"add-transactions/service/config"
	"add-transactions/service/repository"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func SaveRequest(body *string) {
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
		return
	}

	err = repository.CreateItem(av, svc)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %v", err)
		return
	}

	fmt.Println("Successfully added '" + "queryParameter.InstalledPotency" + "' (" + "queryParameter.AnnualProduction" + ") to table " + config.TableName)
}
