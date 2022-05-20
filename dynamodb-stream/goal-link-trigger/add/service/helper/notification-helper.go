package helper

import (
	"add-goal-link/service/i18n"
	"add-goal-link/service/models"
	"add-goal-link/service/repository"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	SkPrefix = "Notification#"
)

// Add a repaid goal notification to the
func AchievedGoalNotification(pk *string, goalName *string, svc *dynamodb.DynamoDB) {
	message := i18n.Notification + *goalName
	notification := models.Notificaion{
		Pk:           pk,
		Notification: &message,
	}
	date := time.Now().Format(time.RFC3339)
	notification.CreationDate = &date
	notification.UpdatedDate = &date
	notification.Sk = SkPrefix + date

	av, err := dynamodbattribute.MarshalMap(notification)
	if err != nil {
		fmt.Printf("Notification item cannot be marshalled: %v", err)
		return
	}
	fmt.Printf("Notification item Marshalled: %+v", av)

	repository.CreateItem(av, svc)
}