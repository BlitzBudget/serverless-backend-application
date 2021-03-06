package helper

import (
	"add-investment-link/service/models"
	"add-investment-link/service/repository"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type InvestmentQueryParameter struct {
	CurrentValue     *float64 `validate:"required" json:":c"`
	InvestmentAmount *float64 `validate:"required" json:":r"`
	UpdatedDate      *string  `validate:"required" json:":u"`
}

// convert item to dynamodb attribute
func InvestmentParseToQueryParameter(currentValue *float64, investmentAmount *float64) map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339Nano)
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
func UpdateCurrentValueForInvestment(svc dynamodbiface.DynamoDBAPI, request *models.Investment) {
	av := InvestmentParseToQueryParameter(request.CurrentValue, request.InvestedAmount)
	updateExpression := "set current_value = :c, invested_amount = :r, updated_date = :u"
	repository.UpdateItem(av, svc, request.Pk, request.Sk, updateExpression)
}
