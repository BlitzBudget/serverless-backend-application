package helper

import (
	"add-goal-link/service/models"
	"add-goal-link/service/repository"
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
func ParseToQueryParameter(goalId *string) map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339)
	av, err := dynamodbattribute.MarshalMap(QueryParameter{
		GoalId:      goalId,
		UpdatedDate: &date,
	})

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to update transaction with goal id %v, %v. \n", *goalId, err)
	}

	return av
}

// Update Transaction Item with Goal ID
func UpdateTransactionWithGoalId(goalId *string, svc *dynamodb.DynamoDB, request *models.Transaction) {
	av := ParseToQueryParameter(goalId)
	updateExpression := "set goal_id = :d, updated_date = :u"
	repository.UpdateItem(av, svc, request.Pk, request.Sk, updateExpression)
}
