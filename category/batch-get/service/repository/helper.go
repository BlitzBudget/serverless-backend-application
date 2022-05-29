package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"get-category/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AttributeBuilder(body *string) (*models.QueryParameter, error) {
	queryParameter := models.QueryParameter{}
	err := json.Unmarshal([]byte(*body), &queryParameter)
	if err != nil {
		fmt.Printf("There was an error marshalling the bytes to struct: %v", err.Error())
		return nil, err
	}

	fmt.Printf("marshalled bytes to struct: %+v", queryParameter)

	return &queryParameter, err
}

func ParseResponse(result *dynamodb.BatchGetItemOutput) (models.ResponseItems, error) {

	if result.Responses == nil {
		msg := "no Items found"
		return nil, errors.New(msg)
	}

	responseItems := models.ResponseItems{}
	var err error

	for _, v := range result.Responses {
		responseItem := models.ResponseItem{}

		for k2, v2 := range v {
			err = dynamodbattribute.UnmarshalMap(v2, &responseItem)
			if err != nil {
				panic(fmt.Sprintf("Failed to unmarshal Record %v, %v", k2, err))
			}
			fmt.Printf("Response Item processed is %v", responseItem.Sk)
			responseItems = append(responseItems, &responseItem)
		}
	}

	fmt.Printf("Parsed %v Items", len(responseItems))
	return responseItems, nil
}
