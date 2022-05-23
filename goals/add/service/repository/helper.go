package repository

import (
	"add-goals/service/config"
	"add-goals/service/models"
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

	goalAchieved := false
	currentAmount := int64(0)
	date := time.Now().Format(time.RFC3339)
	queryParameter.CreationDate = &date
	queryParameter.UpdatedDate = &date
	queryParameter.Sk = config.SkPrefix + date
	if queryParameter.CurrentAmount == nil {
		queryParameter.CurrentAmount = &currentAmount
	}
	if queryParameter.GoalAchieved == nil {
		queryParameter.GoalAchieved = &goalAchieved
	}

	mandatoryFieldsCheck(queryParameter)

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	fmt.Printf("marshalled struct: %+v", av)
	return av, err
}

func mandatoryFieldsCheck(queryParameter models.QueryParameter) {
	if queryParameter.TargetAmount == nil {
		panic(fmt.Sprintln("AttributeBuilder:: Target Amount is empty."))
	}

	if queryParameter.TargetDate == nil {
		panic(fmt.Sprintln("AttributeBuilder:: Target Date is empty."))
	}

	if queryParameter.Name == nil {
		panic(fmt.Sprintln("AttributeBuilder:: Goal Name is empty."))
	}
}
