package helper

import (
	"delete-goal-link/service/models"
	"delete-goal-link/service/repository"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func RemoveGoalLink(records *[]events.DynamoDBEventRecord, svc dynamodbiface.DynamoDBAPI) {

	for _, record := range *records {
		fmt.Printf("RemoveGoalLink:: Old Image is %v. \n", record)
		goal, err := UnmarshalStreamImage(record.Change.OldImage)
		if err != nil {
			fmt.Printf("RemoveGoalLink: Got error unmarshalStreamImage: %v. \n", err)
			continue
		}

		queryOutput, err := repository.GetGoalRuleItems(svc, goal.Pk)
		if err != nil {
			fmt.Printf("RemoveGoalLink: Got error Goal Rule GetGoalRuleItem: %v. \n", err)
			continue
		}

		var goalRules []*models.GoalRule
		goalRules, err = ParseResponse(queryOutput)
		if err != nil {
			fmt.Printf("RemoveGoalLink: Got error Goal Rule ParseResponse: %v. \n", err)
			continue
		}

		for _, v := range goalRules {
			if *v.GoalId == *goal.Sk {
				repository.DeleteItem(v.Pk, v.Sk, svc)
			}
		}

	}
}

func ParseResponse(result *dynamodb.QueryOutput) ([]*models.GoalRule, error) {

	if result.Items == nil {
		msg := "no Items found"
		return nil, errors.New(msg)
	}

	var goalRules []*models.GoalRule
	var err error

	for k, v := range result.Items {
		goalRule := models.GoalRule{}

		err = dynamodbattribute.UnmarshalMap(v, &goalRule)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record %v, %v", k, err))
		}
		goalRules = append(goalRules, &goalRule)
	}

	fmt.Printf("Parsed %v Items. \n", len(goalRules))
	return goalRules, nil
}
