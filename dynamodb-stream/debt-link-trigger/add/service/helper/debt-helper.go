package helper

import (
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Increament the Debt Repayment with the transaction
// Add the transaction amount with the current value field in the debt
func IncrementDebtRepayment(debtRule *models.DebtRule, transaction *models.Transaction, svc *dynamodb.DynamoDB) {
	if !(descriptionMatches(debtRule, transaction.Description) || amountMatches(debtRule, *transaction.Amount)) {
		fmt.Printf("incrementDebtRepayment: The transaction description and the transaction amount do not match: %v. \n", debtRule.Sk)
		return
	}

	debt, err := repository.GetDebtItem(svc, debtRule.DebtId, transaction.Pk)
	if err != nil {
		fmt.Printf("incrementDebtRepayment: Got error Debt GetDebtItem: %v. \n", err)
		return
	}

	fmt.Printf("incrementDebtRepayment: Debt Retireved is : %v", debt.DebtName)
	if *debt.DebtRepaid {
		fmt.Printf("incrementDebtRepayment: The Debt %v has been repaid %v. \n", debt.DebtName, debt.DebtRepaid)
		return
	}

	incrementCurrentValueIfDebtNotRepaid(debt, transaction)
	UpdateCurrentValueForDebt(svc, debt)
}

func incrementCurrentValueIfDebtNotRepaid(debt *models.Debt, transaction *models.Transaction) {
	if *debt.CurrentValue >= *debt.DebtedAmount {
		debtRepaid := true
		debt.DebtRepaid = &debtRepaid
		RepaidDebtNotification(transaction.Pk, debt.DebtName)
	} else {
		incrementDebtCurrentValue(transaction, debt)
	}
}

func incrementDebtCurrentValue(transaction *models.Transaction, debt *models.Debt) {
	currentValue := *transaction.Amount + *debt.CurrentValue
	debt.CurrentValue = &currentValue
	fmt.Printf("incrementDebtRepayment: The new Debt amount is: %v. \n", currentValue)
}
