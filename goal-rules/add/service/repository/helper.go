package repository

import (
	"add-goal-rule/service/config"
	"add-goal-rule/service/models"
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
		fmt.Printf("There was an error marshalling the bytes to struct: %v. \n", err.Error())
		return nil, err
	}

	fmt.Printf("marshalled bytes to struct: %+v. \n", queryParameter)

	date := time.Now().Format(time.RFC3339)
	queryParameter.CreationDate = &date
	queryParameter.UpdatedDate = &date
	queryParameter.Sk = config.SkPrefix + *queryParameter.TransactionName

	mandatoryFieldsCheck(queryParameter)

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	fmt.Printf("marshalled struct: %+v. \n", av)
	return av, err
}


func mandatoryFieldsCheck(queryParameter models.QueryParameter) {
	if queryParameter.GoalId == nil {
		panic(fmt.Sprintln("AttributeBuilder:: Debt Id is empty."))
	}

	if queryParameter.TransactionName == nil {
		panic(fmt.Sprintln("AttributeBuilder:: Transaction Name is empty."))
	}
}
