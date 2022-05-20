package helper

import (
	"add-investment-link/service/models"
	"add-investment-link/service/repository"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Increament the Investment Amount with the transaction
// Add the transaction amount with the current value field in the investment
func IncrementInvestmentAmount(investmentRule *models.InvestmentRule, transaction *models.Transaction, svc *dynamodb.DynamoDB) {
	if !(descriptionMatches(investmentRule, transaction.Description) || amountMatches(investmentRule, *transaction.Amount)) {
		fmt.Printf("incrementInvestmentAmount: The transaction description and the transaction amount do not match: %v. \n", investmentRule.Sk)
		return
	}

	investment, err := repository.GetInvestmentItem(svc, investmentRule.InvestmentId, transaction.Pk)
	if err != nil {
		fmt.Printf("incrementInvestmentAmount: Got error Investment GetInvestmentItem: %v. \n", err)
		return
	}

	fmt.Printf("incrementInvestmentAmount: Investment Retireved is : %v", investment.InvestmentName)

	incrementInvestmentCurrentValue(transaction, investment)
	UpdateCurrentValueForInvestment(svc, investment)
}

func incrementInvestmentCurrentValue(transaction *models.Transaction, investment *models.Investment) {
	currentValue := *transaction.Amount + *investment.CurrentValue
	investedAmount := *transaction.Amount + *investment.InvestmentedAmount
	investment.CurrentValue = &currentValue
	investment.InvestmentedAmount = &investedAmount
	fmt.Printf("incrementInvestmentAmount: The new Investment amount is: %v. \n", currentValue)
}

// Check if the amount matches
func amountMatches(investmentRule *models.InvestmentRule, amt int64) bool {
	return (investmentRule.TransactionAmount != nil && *investmentRule.TransactionAmount == amt) || (investmentRule.TransactionAmount == nil) || *investmentRule.TransactionAmount == int64(0)
}

// Check if the description matches
func descriptionMatches(investmentRule *models.InvestmentRule, description *string) bool {
	return *investmentRule.TransactionName == *description
}
