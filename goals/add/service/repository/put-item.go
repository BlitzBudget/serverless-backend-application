package repository

import (
	"add-goals/service/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateItem(av map[string]*dynamodb.AttributeValue, svc *dynamodb.DynamoDB) error {
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(config.TableName),
	}

	_, err := svc.PutItem(input)
	return err
}
