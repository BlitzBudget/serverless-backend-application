package helper

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type QueryParameter struct {
	GoalId      *string `validate:"required" json:":d"`
	UpdatedDate *string `validate:"required" json:":u"`
}

// convert item to dynamodb attribute
func ParseToQueryParameter() map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339Nano)
	av, err := dynamodbattribute.MarshalMap(QueryParameter{
		GoalId:      nil,
		UpdatedDate: &date,
	})

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to update transaction  %v. \n", err)
	}

	return av
}
