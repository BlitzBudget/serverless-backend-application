package helper

import (
	"add-debt-link/service/config"
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func CreateDebtLink(records *[]events.DynamoDBEventRecord, svc dynamodbiface.DynamoDBAPI) {

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
}

// Fetch Debt
func getDebtRule(description *string, svc dynamodbiface.DynamoDBAPI, pk *string) (*models.DebtRule, error) {
	sk := config.SkDebtRulePrefix + *description
	debtRule, err := repository.GetDebtRuleItem(svc, &sk, pk)
	return debtRule, err
}

// Check if the amount matches
func amountMatches(debtRule *models.DebtRule, amt float64) bool {
	return (debtRule.TransactionAmount != nil && *debtRule.TransactionAmount == amt) || (debtRule.TransactionAmount == nil) || *debtRule.TransactionAmount == float64(0)
}

// Check if the description matches
func descriptionMatches(debtRule *models.DebtRule, description *string) bool {
	return *debtRule.TransactionName == *description
}
