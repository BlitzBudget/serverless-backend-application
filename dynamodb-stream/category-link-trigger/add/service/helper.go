package service

import (
	"add-category-link/service/category-rule/models"
	"add-category-link/service/config"
	"add-category-link/service/repository"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateCategoryLink(records *[]events.DynamoDBEventRecord, svc *dynamodb.DynamoDB) (map[string]*dynamodb.AttributeValue, error) {

	for _, record := range *records {
		description, pk, _, err := unmarshalStreamImage(record.Change.NewImage)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error unmarshalStreamImage: %v", err)
			continue
		}

		sk := config.SkCategoryRulePrefix + *description
		var categoryRule *models.CategoryRule
		categoryRule, err = repository.GetCategoryRuleItem(svc, &sk, pk)
		if err != nil {
			fmt.Printf("AttributeBuilder: Got error Category Rule GetTableItem: %v", err)
			continue
		}

		if categoryRule.TransactionName == description {
			category, err := repository.GetCategoryItem(svc, categoryRule.CategoryId, pk)
			if err != nil {
				fmt.Printf("AttributeBuilder: Got error Category GetTableItem: %v", err)
				continue
			}

			fmt.Printf("AttributeBuilder: Category Retireve is : %v", category.CategoryName)
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
