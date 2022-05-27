package repository

import (
	"fmt"
	"update-category/service/config"
	"update-category/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteItem(av map[string]*dynamodb.AttributeValue, svc *dynamodb.DynamoDB, request *models.RequestModel) {

	input := &dynamodb.DeleteItemInput{
		TableName: &config.TableName,
		Key:       av,
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		panic(fmt.Sprintf("DeleteItem: Failed to update the item %v", err))
	}

	fmt.Println("Deleted the item")
}
