package repository

import (
	getTagModels "add-tag/get-tag/service/models"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func ParseResponse(result *dynamodb.QueryOutput) (getTagModels.ResponseItems, error) {

	if result.Items == nil {
		fmt.Println("ParseResponse:: No Items Found")
		err := fmt.Errorf("ParseResponse:: No Items Found")
		return nil, err
	}

	responseItems := getTagModels.ResponseItems{}
	var err error

	for k, v := range result.Items {
		responseItem := getTagModels.ResponseItem{}

		err = dynamodbattribute.UnmarshalMap(v, &responseItem)
		if err != nil {
			fmt.Printf("Failed to unmarshal Record %v, %v", k, err)
			return nil, err
		}
		responseItems = append(responseItems, &responseItem)
	}

	fmt.Printf("ParseResponse:: Parsed %v Items. \n", len(responseItems))
	return responseItems, nil
}
