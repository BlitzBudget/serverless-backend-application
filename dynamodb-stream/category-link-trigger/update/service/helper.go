package service

import (
	"add-category-link/service/config"
	"add-category-link/service/models"
	"add-category-link/service/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateCategoryLink(records *[]events.DynamoDBEventRecord, svc *dynamodb.DynamoDB) {

	for _, record := range *records {
		transaction, err := UnmarshalStreamImage(record.Change.NewImage)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error unmarshalStreamImage: %v", err)
			continue
		}

		categoryRule, err := getCategoryRule(transaction, svc)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error Category Rule GetTableItem: %v", err)
			continue
		}

		if *categoryRule.TransactionName == *transaction.Description {
			UpdateTransactionWithCategoryId(svc, transaction)
		}
	}
}

func getCategoryRule(transaction *models.Transaction, svc *dynamodb.DynamoDB) (*models.CategoryRule, error) {
	sk := config.SkCategoryRulePrefix + *transaction.Description
	categoryRule, err := repository.GetCategoryRuleItem(svc, &sk, transaction.Pk)
	return categoryRule, err
}
