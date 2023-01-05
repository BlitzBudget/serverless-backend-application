package repository

import (
	"encoding/json"
	"fmt"
	"get-goal/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AttributeBuilder(body *string) (*models.QueryParameter, error) {
	queryParameter := models.QueryParameter{}
	err := json.Unmarshal([]byte(*body), &queryParameter)
	if err != nil {
		fmt.Printf("There was an error marshalling the bytes to struct: %v \n", err.Error())
		return nil, err
	}

	respJSON, _ := json.Marshal(queryParameter)
	fmt.Printf("marshalled bytes to struct: %+v. \n", respJSON)

	return &queryParameter, err
}

func ParseResponse(result *dynamodb.QueryOutput) (models.ResponseItems, error) {

	if result.Items == nil {
		fmt.Println("ParseResponse:: No Items Found")
		err := fmt.Errorf("ParseResponse:: No Items Found")
		return nil, err
	}

	responseItems := models.ResponseItems{}
	var err error

	for k, v := range result.Items {
		responseItem := models.ResponseItem{}

		err = dynamodbattribute.UnmarshalMap(v, &responseItem)
		if err != nil {
			fmt.Printf("Failed to unmarshal Record %v, %v \n", k, err)
			return nil, err
		}
		responseItems = append(responseItems, &responseItem)
	}

	fmt.Printf("Parsed %v Items. \n", len(responseItems))
	return responseItems, nil
}
