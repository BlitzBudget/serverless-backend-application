package service

import (
	"add-debt-link/service/config"
	"add-debt-link/service/models"
	"add-debt-link/service/repository"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateDebtLink(records *[]events.DynamoDBEventRecord, svc *dynamodb.DynamoDB) (map[string]*dynamodb.AttributeValue, error) {

	for _, record := range *records {
		description, pk, amount, err := unmarshalStreamImage(record.Change.NewImage)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error unmarshalStreamImage: %v", err)
			continue
		}

		sk := config.SkDebtRulePrefix + *description
		var debtRule *models.DebtRule
		debtRule, err = repository.GetDebtRuleItem(svc, &sk, pk)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error Debt Rule GetTableItem: %v", err)
			continue
		}

		if debtRule.TransactionName == description {
			debt, err := repository.GetDebtItem(svc, debtRule.DebtId, pk)
			if err != nil {
				fmt.Printf("AttributeBuilder: Got error Debt GetTableItem: %v", err)
				continue
			}

			fmt.Printf("AttributeBuilder: Debt Retireve is : %v", debt.DebtName)
			amt, error := strconv.Atoi(*amount)
			if error != nil {
				fmt.Printf("AttributeBuilder: Unable to convert amount to integer: %v", err)
				continue
			}
			currentValue := int64(amt) + *debt.CurrentValue
			debt.CurrentValue = &currentValue
			fmt.Printf("AttributeBuilder: The new Debt amount is: %v", currentValue)
		}
	}

	return nil, nil
}

// unmarshalStreamImage converts events.DynamoDBAttributeValue to struct
//func unmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {
func unmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue) (*string, *string, *string, error) {

	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range attribute {

		var dbAttr dynamodb.AttributeValue

		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return nil, nil, nil, marshalErr
		}

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr
	}

	fmt.Printf("The Mapping of the new image is %v", dbAttrMap["description"].S)
	// return dynamodbattribute.UnmarshalMap(dbAttrMap, out)
	return dbAttrMap["description"].S, dbAttrMap["pk"].S, dbAttrMap["amount"].N, nil

}
