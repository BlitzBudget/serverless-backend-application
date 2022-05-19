package repository

import (
	"add-category-link/service/category-rule/models"
	"add-category-link/service/config"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// GetTableItem retrieves the item with the year and title from the table
// Inputs:
//     sess is the current session, which provides configuration for the SDK's service clients
//     table is the name of the table
//     title is the movie title
//     year is when the movie was released
// Output:
//     If success, the information about the table item and nil
//     Otherwise, nil and an error from the call to GetItem or UnmarshalMap
func GetCategoryRuleItem(svc dynamodbiface.DynamoDBAPI, sk *string, pk *string) (*models.CategoryRule, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(config.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: pk,
			},
			"sk": {
				S: sk,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		msg := "Could not find '" + *sk + "'"
		return nil, errors.New(msg)
	}

	item := models.CategoryRule{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func GetCategoryItem(svc dynamodbiface.DynamoDBAPI, sk *string, pk *string) (*models.Category, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(config.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: pk,
			},
			"sk": {
				S: sk,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		msg := "Could not find '" + *sk + "'"
		return nil, errors.New(msg)
	}

	item := models.Category{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
