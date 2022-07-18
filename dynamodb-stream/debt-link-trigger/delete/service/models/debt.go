package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create struct to hold info about retrieved item
type Debt struct {
	Pk           *string  `json:"pk"`
	Sk           *string  `json:"sk"`
	CreationDate *string  `json:"creation_date"`
	DebtedAmount *float64 `json:"debted_amount"`
	CurrentValue *float64 `json:"current_value"`
	DebtName     *string  `json:"debt_name"`
	DebtRepaid   *bool    `json:"debt_repaid"`
}

func ConvertDynamoDBToModel(dbAttrMap map[string]*dynamodb.AttributeValue) Debt {

	debt := Debt{
		Pk: dbAttrMap["pk"].S,
		Sk: dbAttrMap["sk"].S,
	}
	return debt
}
