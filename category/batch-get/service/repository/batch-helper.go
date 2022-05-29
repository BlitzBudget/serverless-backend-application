package repository

import (
	"get-category/service/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func BuildAttributes(categoryIds *[]string, userId *string) *dynamodb.BatchGetItemInput {
	mapOfAttrKeys := []map[string]*dynamodb.AttributeValue{}
	tableName := config.TableName

	for _, category := range *categoryIds {
		mapOfAttrKeys = append(mapOfAttrKeys, map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(*userId),
			},
			"sk": {
				S: aws.String(category),
			},
		})
	}

	input := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			tableName: {
				Keys: mapOfAttrKeys,
			},
		},
	}

	return input
}
