package repository

import (
	"add-category-link/service/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func UpdateItem(av map[string]*dynamodb.AttributeValue, svc dynamodbiface.DynamoDBAPI, pk *string, sk *string, updateExpression string) {
	fmt.Printf("Updating the transaction with the sk %v and pk as %v \n", *sk, *pk)

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

	fmt.Printf("UpdateItem: Successfully updated the transaction item %v with category id. \n", *sk)
}
