package repository

import (
	"add-transactions/service/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AttributeBuilder(body *string) (models.QueryParameter, map[string]*dynamodb.AttributeValue, error) {
	queryParameter := models.QueryParameter{}
	err := json.Unmarshal([]byte(*body), &queryParameter)
	if err != nil {
		fmt.Printf("There was an error marshalling the bytes to struct: %v", err.Error())
		return queryParameter, nil, err
	}

	fmt.Printf("marshalled bytes to struct: %+v", queryParameter)

	date := time.Now().Format(time.RFC3339)
	queryParameter.Date = &date

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	fmt.Printf("marshalled struct: %+v", av)
	return queryParameter, av, err
}
