package repository

import (
	"batch-get-category/service/models"
	"encoding/json"
	"errors"
	"fmt"

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

	fmt.Printf("marshalled bytes to struct: %+v. \n", queryParameter)
	return &queryParameter, err
}

func ParseResponse(result *dynamodb.QueryOutput) (models.ResponseItems, error) {

	if result.Items == nil {
		msg := "no Items found"
		fmt.Printf("ParseResponse:: No Items Found")
		return nil, errors.New(msg)
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

func FilterResponse(responseItems []*models.ResponseItem, av *models.QueryParameter) models.ResponseItems {
	var filteredResponseItems models.ResponseItems

	for _, ri := range responseItems {
		if contains(*av.CategoryIds, *ri.Sk) {
			filteredResponseItems = append(filteredResponseItems, ri)
			fmt.Printf("The filtered item is %v. \n", ri.Sk)
		}
	}

	return filteredResponseItems
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
