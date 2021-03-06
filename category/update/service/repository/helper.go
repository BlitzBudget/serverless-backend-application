package repository

import (
	"encoding/json"
	"fmt"
	"time"
	"update-category/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AttributeBuilder(body *string) *models.RequestModel {
	requestModel := models.RequestModel{}
	err := json.Unmarshal([]byte(*body), &requestModel)
	if err != nil {
		panic(fmt.Sprintf("AttributeBuilder: There was an error unmarshalling the bytes to struct: %v", err.Error()))
	}

	fmt.Printf("AttributeBuilder: marshalled bytes to struct: %+v. \n", requestModel)

	return &requestModel
}

func ParseToQueryParameter(request *models.RequestModel) map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339Nano)
	av, err := dynamodbattribute.MarshalMap(models.QueryParameter{
		CategoryName: request.CategoryName,
		CategoryType: request.CategoryType,
		UpdatedDate:  &date,
	})

	if err != nil {
		panic(fmt.Sprintf("ParseToQueryParameter: Failed to marshal request to query parameter data %v", err))
	}

	return av
}
