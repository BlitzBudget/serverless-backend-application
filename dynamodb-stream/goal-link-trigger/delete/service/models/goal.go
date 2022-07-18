package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create struct to hold info about retrieved item
type Goal struct {
	Pk           *string  `json:"pk"`
	Sk           *string  `json:"sk"`
	CreationDate *string  `json:"creation_date"`
	GoaledAmount *float64 `json:"goaled_amount"`
	CurrentValue *float64 `json:"current_value"`
	GoalName     *string  `json:"goal_name"`
	GoalRepaid   *bool    `json:"goal_repaid"`
}

func ConvertDynamoDBToModel(dbAttrMap map[string]*dynamodb.AttributeValue) Goal {

	goal := Goal{
		Pk: dbAttrMap["pk"].S,
		Sk: dbAttrMap["sk"].S,
	}
	return goal
}
