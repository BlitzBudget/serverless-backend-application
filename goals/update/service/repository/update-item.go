package repository

import (
	"fmt"
	"update-goal/service/config"
	"update-goal/service/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func UpdateItem(av map[string]*dynamodb.AttributeValue, svc dynamodbiface.DynamoDBAPI, request *models.RequestModel) error {
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
		UpdateExpression: aws.String("set current_amount = :c, target_amount = :a, target_date = :d, goal_name = :n, goal_achieved = :g, updated_date = :u"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Printf("UpdateItem: Failed to update the item %v \n", err)
		return err
	}

	return nil
}
