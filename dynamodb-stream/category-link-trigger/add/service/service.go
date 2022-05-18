package service

import (
	"add-debt/service/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func ProcessRecords(records *[]events.DynamoDBEventRecord) {
	_, err := repository.AttributeBuilder(records)
	if err != nil {
		panic(fmt.Sprintf("SaveRequest: Got error marshalling new item: %v", err))
	}

	fmt.Println("Successfully added the item!")
}
