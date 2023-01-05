package repository

import (
	"fmt"
	"update-category/service/config"
	"update-category/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteItem(av map[string]*dynamodb.AttributeValue, svc *dynamodb.DynamoDB, request *models.RequestModel) error {

	input := &dynamodb.DeleteItemInput{
		TableName: &config.TableName,
		Key:       av,
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		fmt.Printf("DeleteItem: Failed to update the item %v \n", err)
		return err
	}

	fmt.Println("Deleted the item")
	return nil
}
