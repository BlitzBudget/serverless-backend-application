package helper

import (
	"delete-category-link/service/models"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// unmarshalStreamImage converts events.DynamoDBAttributeValue to struct
//func unmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {
func UnmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue) (*models.Category, error) {

	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range attribute {

		var dbAttr dynamodb.AttributeValue

		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return nil, marshalErr
		}

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr

		fmt.Printf("UnmarshalStreamImage: The parsed DynamoDB attr is %v for the key %v. \n", &dbAttr, k)
	}

	category := models.ConvertDynamoDBToModel(dbAttrMap)

	return &category, nil

}
