package repository

import (
	getTagModels "add-tag/get-tag/service/models"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func ParseResponse(result *dynamodb.QueryOutput) (getTagModels.ResponseItems, error) {

	if result.Items == nil {
		msg := "no Items found"
		return nil, errors.New(msg)
	}

	responseItems := getTagModels.ResponseItems{}
	var err error

	for k, v := range result.Items {
		responseItem := getTagModels.ResponseItem{}

		err = dynamodbattribute.UnmarshalMap(v, &responseItem)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record %v, %v", k, err))
		}
		responseItems = append(responseItems, &responseItem)
	}

	fmt.Printf("Parsed %v Items. \n", len(responseItems))
	return responseItems, nil
}
