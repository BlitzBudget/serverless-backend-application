package repository

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestParseResponse(t *testing.T) {
	dynamoDBAttrValue := map[string]*dynamodb.AttributeValue{
		"Items": {},
	}

	result := dynamodb.QueryOutput{
		Items: []map[string]*dynamodb.AttributeValue{
			dynamoDBAttrValue,
		},
	}

	requestModel, err := ParseResponse(&result)

	if err != nil {
		t.Errorf("Err is null")
		return
	}

	if requestModel == nil {
		t.Errorf("ParseResponse() is null")
		return
	}
}
