package repository

import (
	"add-investment-link/service/config"
	"add-investment-link/service/models"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// GetTableItem retrieves the item with the year and title from the table
// Inputs:
//
//	sess is the current session, which provides configuration for the SDK's service clients
//	table is the name of the table
//	title is the movie title
//	year is when the movie was released
//
// Output:
//
//	If success, the information about the table item and nil
//	Otherwise, nil and an error from the call to GetItem or UnmarshalMap
func GetInvestmentRuleItem(svc dynamodbiface.DynamoDBAPI, sk *string, pk *string) (*models.InvestmentRule, error) {
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
		fmt.Printf("Encountered with an error %v \n", err)
		return nil, err
	}

	if result.Item == nil {
		fmt.Println("The query did not return any items")
		msg := fmt.Sprintf("Could not find the requested item  with sk '%v' and pk as %v", *sk, *pk)
		return nil, errors.New(msg)
	}

	item := models.InvestmentRule{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		fmt.Printf("Encountered with an error %v \n", err)
		return nil, err
	}

	fmt.Printf("Successfully retrieved the item %v. \n", *item.Sk)
	return &item, nil
}

func GetInvestmentItem(svc dynamodbiface.DynamoDBAPI, sk *string, pk *string) (*models.Investment, error) {
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
		fmt.Printf("Encountered with an error %v \n", err)
		return nil, err
	}

	if result.Item == nil {
		fmt.Println("The query did not return any items")
		msg := fmt.Sprintf("Could not find the requested item  with sk '%v' and pk as %v", *sk, *pk)
		return nil, errors.New(msg)
	}

	item := models.Investment{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		fmt.Printf("Encountered with an error %v \n", err)
		return nil, err
	}

	fmt.Printf("Successfully retrieved the item %v. \n", *item.InvestmentName)
	return &item, nil
}
