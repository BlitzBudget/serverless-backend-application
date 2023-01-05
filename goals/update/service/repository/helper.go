package repository

import (
	"encoding/json"
	"fmt"
	"time"
	"update-goal/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AttributeBuilder(body *string) (*models.RequestModel, error) {
	requestModel := models.RequestModel{}
	err := json.Unmarshal([]byte(*body), &requestModel)
	if err != nil {
		fmt.Printf("AttributeBuilder: There was an error unmarshalling the bytes to struct: %v \n", err.Error())
		return nil, err
	}

	fmt.Printf("AttributeBuilder: marshalled bytes to struct: %+v. \n", requestModel)

	return &requestModel, nil
}

func ParseToQueryParameter(request *models.RequestModel) (map[string]*dynamodb.AttributeValue, error) {
	date := time.Now().Format(time.RFC3339Nano)
	av, err := dynamodbattribute.MarshalMap(models.QueryParameter{
		TargetAmount:  request.TargetAmount,
		CurrentAmount: request.CurrentAmount,
		TargetDate:    request.TargetDate,
		GoalAchieved:  request.GoalAchieved,
		Name:          request.Name,
		UpdatedDate:   &date,
	})

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to marshal request to query parameter data %v \n", err)
		return nil, err
	}

	return av, nil
}
