package helper

import (
	"add-goal-link/service/models"
	"add-goal-link/service/repository"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type GoalQueryParameter struct {
	CurrentAmount *float64 `validate:"required" json:":c"`
	GoalAchieved  *bool    `validate:"required" json:":r"`
	UpdatedDate   *string  `validate:"required" json:":u"`
}

// convert item to dynamodb attribute
func GoalParseToQueryParameter(currentValue *float64, goalAchieved *bool) map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339Nano)
	av, err := dynamodbattribute.MarshalMap(GoalQueryParameter{
		CurrentAmount: currentValue,
		GoalAchieved:  goalAchieved,
		UpdatedDate:   &date,
	})

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to update current value %v for the goal %v. \n", *currentValue, err)
	}

	return av
}

// Update Transaction Item with Goal ID
func UpdateCurrentAmountForGoal(svc dynamodbiface.DynamoDBAPI, request *models.Goal) {
	av := GoalParseToQueryParameter(request.CurrentAmount, request.GoalAchieved)
	updateExpression := "set current_amount = :c, goal_achieved = :r, updated_date = :u"
	repository.UpdateItem(av, svc, request.Pk, request.Sk, updateExpression)
}
