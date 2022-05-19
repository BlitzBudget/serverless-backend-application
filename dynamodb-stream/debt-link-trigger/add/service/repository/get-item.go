package repository

import (
	"add-debt-link/service/config"
	"add-debt-link/service/models"
	"errors"
	"fmt"

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
func GetDebtRuleItem(svc dynamodbiface.DynamoDBAPI, sk *string, pk *string) (*models.DebtRule, error) {
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

	item := models.DebtRule{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Successfully retrieved the item %v. \n", item.Sk)
	return &item, nil
}

func GetDebtItem(svc dynamodbiface.DynamoDBAPI, sk *string, pk *string) (*models.Debt, error) {
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

	item := models.Debt{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Successfully retrieved the item %v. \n", item.DebtName)
	return &item, nil
}
