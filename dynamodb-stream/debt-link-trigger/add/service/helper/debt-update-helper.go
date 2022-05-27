package helper

import (
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DebtQueryParameter struct {
	CurrentValue *int64  `validate:"required" json:":c"`
	DebtRepaid   *bool   `validate:"required" json:":r"`
	UpdatedDate  *string `validate:"required" json:":u"`
}

// convert item to dynamodb attribute
func DebtParseToQueryParameter(currentValue *int64, debtRepaid *bool) map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339)
	av, err := dynamodbattribute.MarshalMap(DebtQueryParameter{
		CurrentValue: currentValue,
		DebtRepaid:   debtRepaid,
		UpdatedDate:  &date,
	})

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to update current value %v for the debt %v. \n", *currentValue, err)
	}

	return av
}

// Update Transaction Item with Debt ID
func UpdateCurrentValueForDebt(svc *dynamodb.DynamoDB, request *models.Debt) {
	av := DebtParseToQueryParameter(request.CurrentValue, request.DebtRepaid)
	updateExpression := "set current_value = :c, debt_repaid = :r, updated_date = :u"
	repository.UpdateItem(av, svc, request.Pk, request.Sk, updateExpression)
}
