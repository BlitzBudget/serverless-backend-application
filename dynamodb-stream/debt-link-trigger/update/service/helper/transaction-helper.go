package helper

import (
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type QueryParameter struct {
	DebtId      *string `validate:"required" json:":d"`
	UpdatedDate *string `validate:"required" json:":u"`
}

// convert item to dynamodb attribute
func ParseToQueryParameter(debtId *string) map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339Nano)
	av, err := dynamodbattribute.MarshalMap(QueryParameter{
		DebtId:      debtId,
		UpdatedDate: &date,
	})

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to update transaction with debt id %v, %v. \n", debtId, err)
	}

	return av
}

// Update Transaction Item with Debt ID
func UpdateTransactionWithDebtId(debtId *string, svc *dynamodb.DynamoDB, request *models.Transaction) {
	av := ParseToQueryParameter(debtId)
	updateExpression := "set debt_id = :d, updated_date = :u"
	repository.UpdateItem(av, svc, request.Pk, request.Sk, updateExpression)
}
