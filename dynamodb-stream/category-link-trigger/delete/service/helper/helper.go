package helper

import (
	"delete-category-link/service/config"
	"delete-category-link/service/models"
	"delete-category-link/service/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func RemoveCategoryLink(records *[]events.DynamoDBEventRecord, svc dynamodbiface.DynamoDBAPI) {

	for _, record := range *records {
		fmt.Printf("RemoveGoalLink:: Old Image is %v. \n", record)
		category, err := UnmarshalStreamImage(record.Change.OldImage)
		if err != nil {
			fmt.Printf("RemoveCategoryLink: Got error unmarshalStreamImage: %v. \n", err)
			continue
		}

		wallets := FetchWallets(svc, category)
		for _, wallet := range wallets {
			deleteCategoryRulesForWallet(svc, wallet, category)
		}

	}
}

func deleteCategoryRulesForWallet(svc dynamodbiface.DynamoDBAPI, wallet *models.Wallet, category *models.Category) {
	queryOutput, err := repository.QueryItems(svc, wallet.Sk, config.SkCategoryRulePrefix, config.ProjectionExpression)
	if err != nil {
		fmt.Printf("RemoveCategoryLink: Got error Category Rule GetCategoryRuleItem: %v. \n", err)
		return
	}

	var categoryRules []*models.CategoryRule
	categoryRules, err = ParseResponse(queryOutput)
	if err != nil {
		fmt.Printf("RemoveCategoryLink: Got error Category Rule ParseResponse: %v. \n", err)
		return
	}

	for _, v := range categoryRules {
		if *v.CategoryId == *category.Sk {
			repository.DeleteItem(v.Pk, v.Sk, svc)
		}
	}
}

func ParseResponse(result *dynamodb.QueryOutput) ([]*models.CategoryRule, error) {

	if result.Items == nil {
		fmt.Println("ParseResponse:: No Items Found")
		err := fmt.Errorf("ParseResponse:: No Items Found")
		return nil, err
	}

	var categoryRules []*models.CategoryRule
	var err error

	for k, v := range result.Items {
		categoryRule := models.CategoryRule{}

		err = dynamodbattribute.UnmarshalMap(v, &categoryRule)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record %v, %v \n", k, err))
		}
		categoryRules = append(categoryRules, &categoryRule)
	}

	fmt.Printf("Parsed %v Items. \n", len(categoryRules))
	return categoryRules, nil
}
