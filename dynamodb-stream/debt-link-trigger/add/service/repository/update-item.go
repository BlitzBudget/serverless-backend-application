package repository

import (
	"add-debt-link/service/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func UpdateTransacionItem(av map[string]*dynamodb.AttributeValue, svc *dynamodb.DynamoDB, pk *string, sk *string, updateExpression string) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: av,
		TableName:                 &config.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(*pk),
			},
			"sk": {
				S: aws.String(*sk),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(updateExpression),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Printf("UpdateItem: Failed to update the transaction item %v.\n", err)
	}
	fmt.Printf("UpdateItem: Successfully updated the transaction item %v with debt id. \n", sk)
}
