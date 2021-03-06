package helper

import (
	"add-investment-link/service/config"
	"add-investment-link/service/models"
	"add-investment-link/service/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func CreateInvestmentLink(records *[]events.DynamoDBEventRecord, svc dynamodbiface.DynamoDBAPI) (map[string]*dynamodb.AttributeValue, error) {

	for _, record := range *records {
		transaction, err := UnmarshalStreamImage(record.Change.NewImage)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error unmarshalStreamImage: %v. \n", err)
			continue
		}

		investmentRule, err := getInvestmentRule(transaction.Description, svc, transaction.Pk)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error Investment Rule GetInvestmentRuleItem: %v. \n", err)
			continue
		}

		UpdateTransactionWithInvestmentId(investmentRule.InvestmentId, svc, transaction)
		IncrementInvestmentAmount(investmentRule, transaction, svc)
	}

	return nil, nil
}

// Fetch Investment
func getInvestmentRule(description *string, svc dynamodbiface.DynamoDBAPI, pk *string) (*models.InvestmentRule, error) {
	sk := config.SkInvestmentRulePrefix + *description
	investmentRule, err := repository.GetInvestmentRuleItem(svc, &sk, pk)
	return investmentRule, err
}
