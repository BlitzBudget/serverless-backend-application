package repository

import (
	"add-budget/service/config"
	"add-budget/service/models"
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
	queryParameter.Sk = config.SkPrefix + date
	err = mandatoryFieldsCheck(queryParameter)
	if err != nil {
		return nil, err
	}

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	fmt.Printf("marshalled struct: %+v \n", av)
	return av, err
}

func mandatoryFieldsCheck(queryParameter models.QueryParameter) error {
	if queryParameter.Category == nil {
		fmt.Printf("AttributeBuilder:: Category Id is empty.")
		err := fmt.Errorf("AttributeBuilder:: Category Id is empty")
		return err
	}

	if queryParameter.Planned == nil {
		fmt.Printf("AttributeBuilder:: Planned Amount is empty.")
		err := fmt.Errorf("AttributeBuilder:: Planned Amount is empty")
		return err
	}

	return nil
}
