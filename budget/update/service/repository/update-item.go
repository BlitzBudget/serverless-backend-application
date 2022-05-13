package repository

import (
	"fmt"
	"patch-budget/service/config"
	"patch-budget/service/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func UpdateItem(av map[string]*dynamodb.AttributeValue, svc *dynamodb.DynamoDB, request *models.RequestModel) {
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
		UpdateExpression: aws.String("set amount = :a, description = :d, tags = :t, category = :c, updated_date = :u"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		panic(fmt.Sprintf("UpdateItem: Failed to update the item %v", err))
	}
}
