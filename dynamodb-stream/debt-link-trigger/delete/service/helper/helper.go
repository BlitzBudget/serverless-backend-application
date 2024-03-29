package helper

import (
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func RemoveDebtLink(records *[]events.DynamoDBEventRecord, svc dynamodbiface.DynamoDBAPI) {

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

		var debtRules []*models.DebtRule
		debtRules, err = ParseResponse(queryOutput)
		if err != nil {
			fmt.Printf("RemoveDebtLink: Got error Debt Rule ParseResponse: %v. \n", err)
			continue
		}

		for _, v := range debtRules {
			if *v.DebtId == *debt.Sk {
				repository.DeleteItem(v.Pk, v.Sk, svc)
			}
		}

	}
}

func ParseResponse(result *dynamodb.QueryOutput) ([]*models.DebtRule, error) {

	if result.Items == nil {
		fmt.Println("ParseResponse:: No Items Found")
		err := fmt.Errorf("ParseResponse:: No Items Found")
		return nil, err
	}

	var debtRules []*models.DebtRule
	var err error

	for k, v := range result.Items {
		debtRule := models.DebtRule{}

		err = dynamodbattribute.UnmarshalMap(v, &debtRule)
		if err != nil {
			fmt.Printf("Failed to unmarshal Record %v, %v \n", k, err)
			continue
		}
		debtRules = append(debtRules, &debtRule)
	}

	fmt.Printf("Parsed %v Items. \n", len(debtRules))
	return debtRules, nil
}
