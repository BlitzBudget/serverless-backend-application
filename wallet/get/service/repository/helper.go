package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"get-wallet/service/models"

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

	fmt.Printf("marshalled bytes to struct: %+v. \n", queryParameter)

	return &queryParameter, err
}

func ParseResponse(result *dynamodb.QueryOutput) (models.ResponseItems, error) {

	if result.Items == nil {
		msg := "no Items found"
		return nil, errors.New(msg)
	}

	responseItems := models.ResponseItems{}
	var err error

	for k, v := range result.Items {
		responseItem := models.ResponseItem{}

		err = dynamodbattribute.UnmarshalMap(v, &responseItem)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record %v, %v", k, err))
		}
		responseItems = append(responseItems, &responseItem)
	}

	fmt.Printf("Parsed %v Items. \n", len(responseItems))
	return responseItems, nil
}
