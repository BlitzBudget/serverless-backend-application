package repository

import (
	"add-investment/service/config"
	"add-investment/service/models"
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

	currentValue := float64(0)
	date := time.Now().Format(time.RFC3339Nano)
	queryParameter.CreationDate = &date
	queryParameter.UpdatedDate = &date
	queryParameter.Sk = config.SkPrefix + date

	if queryParameter.CurrentValue == nil {
		queryParameter.CurrentValue = &currentValue
	}
	err = mandatoryFieldsCheck(queryParameter)
	if err != nil {
		fmt.Printf("Encountered with an error %v \n", err)
		return nil, err
	}

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	respJSON, _ = json.Marshal(av)
	fmt.Printf("AttributeBuilder:: marshalled struct: %+v \n", respJSON)
	return av, err
}

func mandatoryFieldsCheck(queryParameter models.QueryParameter) error {
	if queryParameter.InvestedAmount == nil {
		fmt.Println("AttributeBuilder:: Investment Amount is empty.")
		err := fmt.Errorf("AttributeBuilder:: Investment Amount is empty")
		return err
	}

	if queryParameter.InvestmentName == nil {
		fmt.Println("AttributeBuilder:: InvestmentName is empty.")
		err := fmt.Errorf("AttributeBuilder:: InvestmentName is empty")
		return err
	}

	return nil
}
