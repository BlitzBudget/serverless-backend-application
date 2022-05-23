package repository

import (
	"add-category/service/config"
	"add-category/service/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AttributeBuilder(body *string) (map[string]*dynamodb.AttributeValue, error) {
	queryParameter := models.QueryParameter{}
	err := json.Unmarshal([]byte(*body), &queryParameter)
	if err != nil {
		fmt.Printf("There was an error marshalling the bytes to struct: %v", err.Error())
		return nil, err
	}

	fmt.Printf("marshalled bytes to struct: %+v", queryParameter)

	date := time.Now().Format(time.RFC3339)
	queryParameter.CreationDate = &date
	queryParameter.UpdatedDate = &date
	queryParameter.Sk = config.SkPrefix + date

	mandatoryFieldsCheck(queryParameter)

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	fmt.Printf("marshalled struct: %+v", av)
	return av, err
}

func mandatoryFieldsCheck(queryParameter models.QueryParameter) {
	if queryParameter.CategoryName == nil {
		panic(fmt.Sprintln("AttributeBuilder:: Category Name is empty."))
	}

	if queryParameter.CategoryType == nil {
		panic(fmt.Sprintln("AttributeBuilder:: Category Type is empty."))
	}
}
