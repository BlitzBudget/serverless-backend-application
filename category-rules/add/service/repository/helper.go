package repository

import (
	"add-category-rule/service/config"
	"add-category-rule/service/models"
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
		fmt.Printf("There was an error marshalling the bytes to struct: %v \n", err.Error())
		return nil, err
	}

	fmt.Printf("marshalled bytes to struct: %+v. \n", queryParameter)

	date := time.Now().Format(time.RFC3339Nano)
	queryParameter.CreationDate = &date
	queryParameter.UpdatedDate = &date
	queryParameter.Sk = config.SkPrefix + *queryParameter.TransactionName
	err = mandatoryFieldsCheck(queryParameter)
	if err != nil {
		return nil, err
	}

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	fmt.Printf("marshalled struct: %+v \n", av)
	return av, err
}

func mandatoryFieldsCheck(queryParameter models.QueryParameter) error {
	if queryParameter.CategoryId == nil {
		fmt.Println("AttributeBuilder:: Category Id is empty.")
		err := fmt.Errorf("AttributeBuilder:: Category Id is empty")
		return err
	}

	if queryParameter.TransactionName == nil {
		fmt.Println("AttributeBuilder:: Transaction Name is empty.")
		err := fmt.Errorf("AttributeBuilder:: Transaction Name is empty")
		return err
	}

	return nil
}
