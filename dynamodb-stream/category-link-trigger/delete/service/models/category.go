package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create struct to hold info about retrieved item
type Category struct {
	Pk               *string  `json:"pk"`
	Sk               *string  `json:"sk"`
	CreationDate     *string  `json:"creation_date"`
	CategoryedAmount *float64 `json:"categoryed_amount"`
	CurrentValue     *float64 `json:"current_value"`
	CategoryName     *string  `json:"category_name"`
	CategoryRepaid   *bool    `json:"category_repaid"`
}

func ConvertDynamoDBToModel(dbAttrMap map[string]*dynamodb.AttributeValue) Category {

	category := Category{
		Pk: dbAttrMap["pk"].S,
		Sk: dbAttrMap["sk"].S,
	}
	return category
}
