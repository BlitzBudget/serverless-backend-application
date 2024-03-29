package service

import (
	"add-category-link/service/config"
	"add-category-link/service/models"
	"add-category-link/service/repository"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func CreateCategoryLink(records *[]events.DynamoDBEventRecord, svc dynamodbiface.DynamoDBAPI) {

	for _, record := range *records {
		transaction, err := UnmarshalStreamImage(record.Change.NewImage)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error unmarshalStreamImage: %v \n", err)
			continue
		}

		categoryRule, err := getCategoryRule(transaction, svc)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error Category Rule GetTableItem: %v \n", err)
			continue
		}

		if strings.Contains(*transaction.Description, *categoryRule.TransactionName) {
			UpdateTransactionWithCategoryId(categoryRule.CategoryId, svc, transaction)
		}
	}
}

func getCategoryRule(transaction *models.Transaction, svc dynamodbiface.DynamoDBAPI) (*models.CategoryRule, error) {
	sk := config.SkCategoryRulePrefix + *transaction.Description
	categoryRule, err := repository.GetCategoryRuleItem(svc, &sk, transaction.Pk)
	return categoryRule, err
}
