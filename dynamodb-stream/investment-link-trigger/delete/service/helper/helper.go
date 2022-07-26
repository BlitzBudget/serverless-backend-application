package helper

import (
	"add-investment-link/service/models"
	"add-investment-link/service/repository"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func RemoveInvestmentLink(records *[]events.DynamoDBEventRecord, svc dynamodbiface.DynamoDBAPI) {

	for _, record := range *records {
		fmt.Printf("RemoveGoalLink:: Old Image is %v. \n", record)
		investment, err := UnmarshalStreamImage(record.Change.OldImage)
		if err != nil {
			fmt.Printf("RemoveInvestmentLink: Got error unmarshalStreamImage: %v. \n", err)
			continue
		}

		queryOutput, err := repository.GetInvestmentRuleItems(svc, investment.Pk)
		if err != nil {
			fmt.Printf("RemoveInvestmentLink: Got error Investment Rule GetInvestmentRuleItem: %v. \n", err)
			continue
		}

		var investmentRules []*models.InvestmentRule
		investmentRules, err = ParseResponse(queryOutput)
		if err != nil {
			fmt.Printf("RemoveInvestmentLink: Got error Investment Rule ParseResponse: %v. \n", err)
			continue
		}

		for _, v := range investmentRules {
			if *v.InvestmentId == *investment.Sk {
				repository.DeleteItem(v.Pk, v.Sk, svc)
			}
		}

	}
}

func ParseResponse(result *dynamodb.QueryOutput) ([]*models.InvestmentRule, error) {

	if result.Items == nil {
		msg := "no Items found"
		return nil, errors.New(msg)
	}

	var investmentRules []*models.InvestmentRule
	var err error

	for k, v := range result.Items {
		investmentRule := models.InvestmentRule{}

		err = dynamodbattribute.UnmarshalMap(v, &investmentRule)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record %v, %v", k, err))
		}
		investmentRules = append(investmentRules, &investmentRule)
	}

	fmt.Printf("Parsed %v Items. \n", len(investmentRules))
	return investmentRules, nil
}
