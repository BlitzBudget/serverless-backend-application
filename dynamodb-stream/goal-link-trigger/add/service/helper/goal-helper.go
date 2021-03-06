package helper

import (
	"add-goal-link/service/models"
	"add-goal-link/service/repository"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Increament the Goal Achieved with the transaction
// Add the transaction amount with the current value field in the goal
func IncrementGoalAchieved(goalRule *models.GoalRule, transaction *models.Transaction, svc dynamodbiface.DynamoDBAPI) {
	if !(descriptionMatches(goalRule, transaction.Description) || amountMatches(goalRule, *transaction.Amount)) {
		fmt.Printf("incrementGoalAchieved: The transaction description and the transaction amount do not match: %v. \n", *goalRule.Sk)
		return
	}

	goal, err := repository.GetGoalItem(svc, goalRule.GoalId, transaction.Pk)
	if err != nil {
		fmt.Printf("incrementGoalAchieved: Got error Goal GetGoalItem: %v. \n", err)
		return
	}

	fmt.Printf("incrementGoalAchieved: Goal Retireved is : %v. \n", goal.GoalName)
	if *goal.GoalAchieved {
		fmt.Printf("incrementGoalAchieved: The Goal %v has been repaid %v. \n", goal.GoalName, goal.GoalAchieved)
		return
	}

	incrementGoalCurrentAmount(transaction, goal)
	updategoalAchieved(goal, transaction, svc)
	UpdateCurrentAmountForGoal(svc, goal)
}

func updategoalAchieved(goal *models.Goal, transaction *models.Transaction, svc dynamodbiface.DynamoDBAPI) {
	if *goal.CurrentAmount >= *goal.TargetAmount {
		goalAchieved := true
		goal.GoalAchieved = &goalAchieved
		AchievedGoalNotification(transaction.Pk, goal.GoalName, svc)
	}
}

func incrementGoalCurrentAmount(transaction *models.Transaction, goal *models.Goal) {
	currentValue := *transaction.Amount + *goal.CurrentAmount
	goal.CurrentAmount = &currentValue
	fmt.Printf("incrementGoalAchieved: The new Goal amount is: %v. \n", currentValue)
}
