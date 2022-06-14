package helper

import (
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func RemoveDebtLink(records *[]events.DynamoDBEventRecord, svc *dynamodb.DynamoDB) {

	for _, record := range *records {
		fmt.Printf("RemoveGoalLink:: Old Image is %v. \n", record)
		debt, err := UnmarshalStreamImage(record.Change.OldImage)
		if err != nil {
			fmt.Printf("RemoveDebtLink: Got error unmarshalStreamImage: %v. \n", err)
			continue
		}

		queryOutput, err := repository.GetDebtRuleItems(svc, debt.Pk)
		if err != nil {
			fmt.Printf("RemoveDebtLink: Got error Debt Rule GetDebtRuleItem: %v. \n", err)
			continue
		}

		var debtRules models.DebtRules
		debtRules, err = ParseResponse(queryOutput)
		if err != nil {
			fmt.Printf("RemoveDebtLink: Got error Debt Rule ParseResponse: %v. \n", err)
			continue
		}

		for _, v := range debtRules {
			if v.Sk == debt.Sk {
				repository.DeleteItem(v.Pk, v.Sk, svc)
			}
		}

		fmt.Printf("Successfully deleted the Debt rules for Debt %v", debt.Sk)

	}
}

func ParseResponse(result *dynamodb.QueryOutput) (models.DebtRules, error) {

	if result.Items == nil {
		msg := "no Items found"
		return nil, errors.New(msg)
	}

	debtRules := models.DebtRules{}
	var err error

	for k, v := range result.Items {
		debtRule := models.DebtRule{}

		err = dynamodbattribute.UnmarshalMap(v, &debtRule)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record %v, %v", k, err))
		}
		debtRules = append(debtRules, &debtRule)
	}

	fmt.Printf("Parsed %v Items", len(debtRules))
	return debtRules, nil
}
