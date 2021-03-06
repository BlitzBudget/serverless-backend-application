package helper

import (
	"add-goal-link/service/config"
	"add-goal-link/service/models"
	"add-goal-link/service/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func CreateGoalLink(records *[]events.DynamoDBEventRecord, svc dynamodbiface.DynamoDBAPI) {

	for _, record := range *records {
		transaction, err := UnmarshalStreamImage(record.Change.NewImage)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error unmarshalStreamImage: %v. \n", err)
			continue
		}

		goalRule, err := getGoalRule(transaction.Description, svc, transaction.Pk)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error Goal Rule GetGoalRuleItem: %v. \n", err)
			continue
		}

		UpdateTransactionWithGoalId(svc, transaction)
		IncrementGoalAchieved(goalRule, transaction, svc)
	}
}

// Fetch Goal
func getGoalRule(description *string, svc dynamodbiface.DynamoDBAPI, pk *string) (*models.GoalRule, error) {
	sk := config.SkGoalRulePrefix + *description
	goalRule, err := repository.GetGoalRuleItem(svc, &sk, pk)
	return goalRule, err
}

// Check if the amount matches
func amountMatches(goalRule *models.GoalRule, amt float64) bool {
	return (goalRule.TransactionAmount != nil && *goalRule.TransactionAmount == amt) || (goalRule.TransactionAmount == nil) || *goalRule.TransactionAmount == float64(0)
}

// Check if the description matches
func descriptionMatches(goalRule *models.GoalRule, description *string) bool {
	return *goalRule.TransactionName == *description
}
