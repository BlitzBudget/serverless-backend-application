package helper

import (
	"add-investment-link/service/models"
	"add-investment-link/service/repository"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type QueryParameter struct {
	InvestmentId *string `validate:"required" json:":d"`
	UpdatedDate  *string `validate:"required" json:":u"`
}

// convert item to dynamodb attribute
func ParseToQueryParameter(investmentId *string) map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339)
	av, err := dynamodbattribute.MarshalMap(QueryParameter{
		InvestmentId: investmentId,
		UpdatedDate:  &date,
	})

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to update transaction with investment id %v, %v. \n", *investmentId, err)
	}

	return av
}

// Update Transaction Item with Investment ID
func UpdateTransactionWithInvestmentId(investmentId *string, svc *dynamodb.DynamoDB, request *models.Transaction) {
	av := ParseToQueryParameter(investmentId)
	updateExpression := "set investment_id = :d, updated_date = :u"
	repository.UpdateItem(av, svc, request.Pk, request.Sk, updateExpression)
}
