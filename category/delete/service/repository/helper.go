package repository

import (
	"encoding/json"
	"fmt"
	"update-category/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AttributeBuilder(body *string) (*models.RequestModel, error) {
	requestModel := models.RequestModel{}
	err := json.Unmarshal([]byte(*body), &requestModel)
	if err != nil {
		fmt.Printf("AttributeBuilder: There was an error unmarshalling the bytes to struct: %v", err.Error())
		return nil, err
	}

	fmt.Printf("AttributeBuilder: marshalled bytes to struct: %+v. \n", requestModel)

	return &requestModel, nil
}

func ParseToQueryParameter(request *models.RequestModel) (map[string]*dynamodb.AttributeValue, error) {
	av, err := dynamodbattribute.MarshalMap(request)

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to marshal request to query parameter data %v", err)
		return nil, err
	}

	return av, nil
}
