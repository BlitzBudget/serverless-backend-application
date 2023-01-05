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
		fmt.Printf("There was an error marshalling the bytes to struct: %v \n", err.Error())
		return nil, err
	}

	respJSON, _ := json.Marshal(queryParameter)
	fmt.Printf("marshalled bytes to struct: %+v. \n", respJSON)

	goalAchieved := false
	currentAmount := float64(0)
	date := time.Now().Format(time.RFC3339Nano)
	queryParameter.CreationDate = &date
	queryParameter.UpdatedDate = &date
	queryParameter.Sk = config.SkPrefix + date
	if queryParameter.CurrentAmount == nil {
		queryParameter.CurrentAmount = &currentAmount
	}
	if queryParameter.GoalAchieved == nil {
		queryParameter.GoalAchieved = &goalAchieved
	}
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
	if queryParameter.TargetAmount == nil {
		fmt.Println("AttributeBuilder:: Target Amount is empty.")
		err := fmt.Errorf("AttributeBuilder:: Target Amount is empty")
		return err
	}

	if queryParameter.TargetDate == nil {
		fmt.Sprintln("AttributeBuilder:: Target Date is empty.")
		err := fmt.Errorf("AttributeBuilder:: Target Date is empty")
		return err
	}

	if queryParameter.Name == nil {
		fmt.Println("AttributeBuilder:: Goal Name is empty.")
		err := fmt.Errorf("AttributeBuilder:: Goal Name is empty")
		return err
	}

	return nil
}
