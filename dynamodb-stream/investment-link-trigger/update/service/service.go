package service

import (
	"add-investment-link/service/helper"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ProcessRecords(records *[]events.DynamoDBEventRecord) {
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
	fmt.Printf("Records for processing %v. \n", records)

	_, err := helper.CreateInvestmentLink(records, svc)
	if err != nil {
		panic(fmt.Sprintf("SaveRequest: Got error marshalling new item: %v \n", err))
	}

	fmt.Println("Successfully processed all the records!")
}
