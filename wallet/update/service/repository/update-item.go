package repository

import (
	"fmt"
	"update-wallet/service/config"
	"update-wallet/service/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func UpdateItem(av map[string]*dynamodb.AttributeValue, svc dynamodbiface.DynamoDBAPI, request *models.RequestModel) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: av,
		TableName:                 &config.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(*request.Pk),
			},
			"sk": {
				S: aws.String(*request.Sk),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set wallet_currency = :c, wallet_name = :n, updated_date = :u"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		panic(fmt.Sprintf("UpdateItem: Failed to update the item %v", err))
	}
}
