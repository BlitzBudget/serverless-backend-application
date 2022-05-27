package helper

import (
	"add-debt-link/service/i18n"
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	SkPrefix = "Notification#"
)

// Add a repaid debt notification to the
func RepaidDebtNotification(pk *string, debtName *string, svc *dynamodb.DynamoDB) {
	message := i18n.Notification + *debtName
	notification := models.Notificaion{
		Pk:           pk,
		Notification: &message,
	}
	date := time.Now().Format(time.RFC3339Nano)
	notification.CreationDate = &date
	notification.UpdatedDate = &date
	notification.Sk = SkPrefix + message

	av, err := dynamodbattribute.MarshalMap(notification)
	if err != nil {
		fmt.Printf("Notification item cannot be marshalled: %v", err)
		return
	}
	fmt.Printf("Notification item Marshalled: %+v", av)

	repository.CreateItem(av, svc)
}
