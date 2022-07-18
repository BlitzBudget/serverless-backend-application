package helper

import (
	"add-debt-link/service/config"
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func RemoveDebtLink(records *[]events.DynamoDBEventRecord, svc *dynamodb.DynamoDB) {

	for _, record := range *records {
		fmt.Printf("RemoveGoalLink:: Old Image is %v. \n", record)
		transaction, err := UnmarshalStreamImage(record.Change.OldImage)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error unmarshalStreamImage: %v. \n", err)
			continue
		}

		debtRule, err := getDebtRule(transaction.Description, svc, transaction.Pk)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error Debt Rule GetDebtRuleItem: %v. \n", err)
			continue
		}

		IncrementDebtRepayment(debtRule, transaction, svc)
	}
}

// Fetch Debt
func getDebtRule(description *string, svc *dynamodb.DynamoDB, pk *string) (*models.DebtRule, error) {
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
