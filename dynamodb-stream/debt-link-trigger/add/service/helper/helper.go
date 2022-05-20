package helper

import (
	"add-debt-link/service/config"
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateDebtLink(records *[]events.DynamoDBEventRecord, svc *dynamodb.DynamoDB) (map[string]*dynamodb.AttributeValue, error) {

	for _, record := range *records {
		transaction, err := UnmarshalStreamImage(record.Change.NewImage)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error unmarshalStreamImage: %v. \n", err)
			continue
		}

		debtRule, err := getDebtRule(transaction.Description, svc, transaction.Pk)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error Debt Rule GetDebtRuleItem: %v. \n", err)
			continue
		}

		UpdateTransactionWithDebtId(debtRule.DebtId, svc, transaction)
		IncrementDebtRepayment(debtRule, transaction, svc)
	}

	return nil, nil
}

// Fetch Debt
func getDebtRule(description *string, svc *dynamodb.DynamoDB, pk *string) (*models.DebtRule, error) {
	sk := config.SkDebtRulePrefix + *description
	debtRule, err := repository.GetDebtRuleItem(svc, &sk, pk)
	return debtRule, err
}

// Check if the amount matches
func amountMatches(debtRule *models.DebtRule, amt int64) bool {
	return (debtRule.TransactionAmount != nil && *debtRule.TransactionAmount == amt) || (debtRule.TransactionAmount == nil) || *debtRule.TransactionAmount == int64(0)
}

// Check if the description matches
func descriptionMatches(debtRule *models.DebtRule, description *string) bool {
	return *debtRule.TransactionName == *description
}
