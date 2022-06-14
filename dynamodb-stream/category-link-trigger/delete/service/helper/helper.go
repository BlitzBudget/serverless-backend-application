package helper

import (
	"delete-category-link/service/config"
	"delete-category-link/service/models"
	"delete-category-link/service/repository"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func RemoveCategoryLink(records *[]events.DynamoDBEventRecord, svc *dynamodb.DynamoDB) {

	for _, record := range *records {
		fmt.Printf("RemoveGoalLink:: Old Image is %v. \n", record)
		category, err := UnmarshalStreamImage(record.Change.OldImage)
		if err != nil {
			fmt.Printf("RemoveCategoryLink: Got error unmarshalStreamImage: %v. \n", err)
			continue
		}

		fetchWallets(svc, category, err)

		queryOutput, err := repository.QueryItems(svc, category.Pk, config.SkCategoryRulePrefix)
		if err != nil {
			fmt.Printf("RemoveCategoryLink: Got error Category Rule GetCategoryRuleItem: %v. \n", err)
			continue
		}

		var categoryRules []*models.CategoryRule
		categoryRules, err = ParseResponse(queryOutput)
		if err != nil {
			fmt.Printf("RemoveCategoryLink: Got error Category Rule ParseResponse: %v. \n", err)
			continue
		}

		for _, v := range categoryRules {
			if *v.CategoryId == *category.Sk {
				repository.DeleteItem(v.Pk, v.Sk, svc)
			}
		}

	}
}

func fetchWallets(svc *dynamodb.DynamoDB, category *models.Category, err error) {
	// TODO
}

func ParseResponse(result *dynamodb.QueryOutput) ([]*models.CategoryRule, error) {

	if result.Items == nil {
		msg := "no Items found"
		return nil, errors.New(msg)
	}

	var categoryRules []*models.CategoryRule
	var err error

	for k, v := range result.Items {
		categoryRule := models.CategoryRule{}

		err = dynamodbattribute.UnmarshalMap(v, &categoryRule)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record %v, %v", k, err))
		}
		categoryRules = append(categoryRules, &categoryRule)
	}

	fmt.Printf("Parsed %v Items. \n", len(categoryRules))
	return categoryRules, nil
}
