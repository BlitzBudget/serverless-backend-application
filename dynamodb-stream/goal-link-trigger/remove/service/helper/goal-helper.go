package helper

import (
	"add-goal-link/service/models"
	"add-goal-link/service/repository"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Increament the Goal Achieved with the transaction
// Add the transaction amount with the current value field in the goal
func IncrementGoalAchieved(goalRule *models.GoalRule, transaction *models.Transaction, svc *dynamodb.DynamoDB) {
	if !(descriptionMatches(goalRule, transaction.Description) || amountMatches(goalRule, *transaction.Amount)) {
		fmt.Printf("incrementGoalAchieved: The transaction description and the transaction amount do not match: %v. \n", goalRule.Sk)
		return
	}

	goal, err := repository.GetGoalItem(svc, goalRule.GoalId, transaction.Pk)
	if err != nil {
		fmt.Printf("incrementGoalAchieved: Got error Goal GetGoalItem: %v. \n", err)
		return
	}

	fmt.Printf("incrementGoalAchieved: Goal Retireved is : %v. \n", goal.GoalName)

	incrementGoalCurrentAmount(transaction, goal)
	updategoalAchieved(goal, transaction, svc)
	UpdateCurrentAmountForGoal(svc, goal)
}

func updategoalAchieved(goal *models.Goal, transaction *models.Transaction, svc *dynamodb.DynamoDB) {
	if *goal.CurrentAmount <= *goal.TargetAmount {
		goalAchieved := false
		goal.GoalAchieved = &goalAchieved
	}
}

func incrementGoalCurrentAmount(transaction *models.Transaction, goal *models.Goal) {
	currentValue := *goal.CurrentAmount - *transaction.Amount
	goal.CurrentAmount = &currentValue
	fmt.Printf("incrementGoalAchieved: The new Goal amount is: %v. \n", currentValue)
}
