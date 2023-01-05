package repository

import (
	"add-investment-rule/service/config"
	"add-investment-rule/service/models"
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

	respJSON, _ := json.Marshal(queryParameter)
	fmt.Printf("marshalled bytes to struct: %+v. \n", respJSON)

	date := time.Now().Format(time.RFC3339Nano)
	queryParameter.CreationDate = &date
	queryParameter.UpdatedDate = &date
	queryParameter.Sk = config.SkPrefix + *queryParameter.TransactionName
	err = mandatoryFieldsCheck(queryParameter)
	if err != nil {
		return nil, err
	}

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	respJSON, _ := json.Marshal(av)
	fmt.Printf("AttributeBuilder:: marshalled struct: %+v \n", respJSON)
	return av, err
}

func mandatoryFieldsCheck(queryParameter models.QueryParameter) error {
	if queryParameter.InvestmentId == nil {
		fmt.Println("AttributeBuilder:: Investment Id is empty.")
		err := fmt.Errorf("AttributeBuilder:: Investment Id is empty")
		return err
	}

	if queryParameter.TransactionName == nil {
		fmt.Println("AttributeBuilder:: Transaction Name is empty.")
		err := fmt.Errorf("AttributeBuilder:: Transaction Name is empty")
		return err
	}

	return nil
}
