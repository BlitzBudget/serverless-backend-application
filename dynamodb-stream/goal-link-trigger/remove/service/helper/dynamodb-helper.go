package helper

import (
	"add-goal-link/service/models"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// unmarshalStreamImage converts events.DynamoDBAttributeValue to struct
//func unmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {
func UnmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue) (*models.Transaction, error) {

	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range attribute {
		var dbAttr dynamodb.AttributeValue

		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return nil, marshalErr
		}

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr
		fmt.Printf("Mapping attribute %v with Key %v. \n", dbAttr, k)
	}

	transaction := models.ConvertDynamoDBToModel(dbAttrMap)

	return &transaction, nil

}
