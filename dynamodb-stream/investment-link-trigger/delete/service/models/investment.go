package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create struct to hold info about retrieved item
type Investment struct {
	Pk             *string  `json:"pk"`
	Sk             *string  `json:"sk"`
	CreationDate   *string  `json:"creation_date"`
	InvestedAmount *float64 `json:"invested_amount"`
	CurrentValue   *float64 `json:"current_value"`
	InvestmentName *string  `json:"investment_name"`
}

func ConvertDynamoDBToModel(dbAttrMap map[string]*dynamodb.AttributeValue) Investment {

	investment := Investment{
		Pk: dbAttrMap["pk"].S,
		Sk: dbAttrMap["sk"].S,
	}
	return investment
}
