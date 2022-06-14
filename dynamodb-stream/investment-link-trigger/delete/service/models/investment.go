package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create struct to hold info about retrieved item
type Investment struct {
	Pk                 *string `json:"pk"`
	Sk                 *string `json:"sk"`
	CreationDate       *string `json:"creation_date"`
	InvestmentedAmount *int64  `json:"investmented_amount"`
	CurrentValue       *int64  `json:"current_value"`
	InvestmentName     *string `json:"investment_name"`
	InvestmentRepaid   *bool   `json:"investment_repaid"`
}

func ConvertDynamoDBToModel(dbAttrMap map[string]*dynamodb.AttributeValue) Investment {

	investment := Investment{
		Pk: dbAttrMap["pk"].S,
		Sk: dbAttrMap["sk"].S,
	}
	return investment
}
