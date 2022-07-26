package repository

import (
	"add-debt-link/service/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func CreateItem(av map[string]*dynamodb.AttributeValue, svc dynamodbiface.DynamoDBAPI) error {
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(config.TableName),
	}

	_, err := svc.PutItem(input)
	fmt.Printf("Add Item: Successfully added the notification item for successful debt repayment. \n")
	return err
}
