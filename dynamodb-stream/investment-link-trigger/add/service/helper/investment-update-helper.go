package helper

import (
	"add-investment-link/service/models"
	"add-investment-link/service/repository"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type InvestmentQueryParameter struct {
	CurrentValue     *int64  `validate:"required" json:":c"`
	InvestmentAmount *bool   `validate:"required" json:":r"`
	UpdatedDate      *string `validate:"required" json:":u"`
}

// convert item to dynamodb attribute
func InvestmentParseToQueryParameter(currentValue *int64, investmentAmount *bool) map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339)
	av, err := dynamodbattribute.MarshalMap(InvestmentQueryParameter{
		CurrentValue:     currentValue,
		InvestmentAmount: investmentAmount,
		UpdatedDate:      &date,
	})

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to update current value %v for the investment %v. \n", *currentValue, err)
	}

	return av
}

// Update Transaction Item with Investment ID
func UpdateCurrentValueForInvestment(svc *dynamodb.DynamoDB, request *models.Investment) {
	av := InvestmentParseToQueryParameter(request.CurrentValue, request.InvestmentAmount)
	updateExpression := "set current_value = :c, investment_amount = :r, updated_date = :u"
	repository.UpdateItem(av, svc, request.Pk, request.Sk, updateExpression)
}
