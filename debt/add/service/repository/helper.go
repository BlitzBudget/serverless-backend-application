package repository

import (
	"add-debt/service/config"
	"add-debt/service/models"
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

	debtRepaid := false
	date := time.Now().Format(time.RFC3339Nano)
	queryParameter.CreationDate = &date
	queryParameter.UpdatedDate = &date
	queryParameter.Sk = config.SkPrefix + date

	mandatoryFieldsCheck(queryParameter)

	if queryParameter.DebtRepaid == nil {
		queryParameter.DebtRepaid = &debtRepaid
	}

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	fmt.Printf("marshalled struct: %+v", av)
	return av, err
}

func mandatoryFieldsCheck(queryParameter models.QueryParameter) {
	if queryParameter.DebtedAmount == nil {
		panic(fmt.Sprintln("AttributeBuilder:: Debt Amount is empty."))
	}

	if queryParameter.DebtName == nil {
		panic(fmt.Sprintln("AttributeBuilder:: Debt Name is empty."))
	}
}
