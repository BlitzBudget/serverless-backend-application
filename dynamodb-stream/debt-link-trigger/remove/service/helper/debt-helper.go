package helper

import (
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Increament the Debt Repayment with the transaction
// Add the transaction amount with the current value field in the debt
func IncrementDebtRepayment(debtRule *models.DebtRule, transaction *models.Transaction, svc dynamodbiface.DynamoDBAPI) {
	if !(descriptionMatches(debtRule, transaction.Description) || amountMatches(debtRule, *transaction.Amount)) {
		fmt.Printf("incrementDebtRepayment: The transaction description and the transaction amount do not match: %v. \n", debtRule.Sk)
		return
	}

	debt, err := repository.GetDebtItem(svc, debtRule.DebtId, transaction.Pk)
	if err != nil {
		fmt.Printf("incrementDebtRepayment: Got error Debt GetDebtItem: %v. \n", err)
		return
	}

	fmt.Printf("incrementDebtRepayment: Debt Retireved is : %v. \n", debt.DebtName)

	incrementDebtCurrentValue(transaction, debt)
	updateDebtRepaid(debt, transaction, svc)
	UpdateCurrentValueForDebt(svc, debt)
}

func updateDebtRepaid(debt *models.Debt, transaction *models.Transaction, svc dynamodbiface.DynamoDBAPI) {
	if *debt.CurrentValue <= *debt.DebtedAmount {
		debtRepaid := false
		debt.DebtRepaid = &debtRepaid
	}
}

func incrementDebtCurrentValue(transaction *models.Transaction, debt *models.Debt) {
	currentValue := *debt.CurrentValue - *transaction.Amount
	debt.CurrentValue = &currentValue
	fmt.Printf("incrementDebtRepayment: The new Debt amount is: %v. \n", currentValue)
}
