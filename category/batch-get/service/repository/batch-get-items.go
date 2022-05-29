package repository

import (
	"get-category/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func BatchGetItems(av *models.QueryParameter, svc *dynamodb.DynamoDB) (*dynamodb.BatchGetItemOutput, error) {
	input := BuildAttributes(av.CategoryIDs, av.UserId)

	data, err := svc.BatchGetItem(input)
	return data, err
}
