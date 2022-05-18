package repository

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func AttributeBuilder(records *[]events.DynamoDBEventRecord) (map[string]*dynamodb.AttributeValue, error) {

	for _, record := range *records {
		unmarshalStreamImage(record.Change.NewImage)
	}

	return nil, nil
}

// unmarshalStreamImage converts events.DynamoDBAttributeValue to struct
//func unmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {
func unmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue) error {

	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range attribute {

		var dbAttr dynamodb.AttributeValue

		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return marshalErr
		}

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr
	}

	fmt.Printf("The Mapping of the new image is %v", dbAttrMap)
	// return dynamodbattribute.UnmarshalMap(dbAttrMap, out)
	return nil

}
