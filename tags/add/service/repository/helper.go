package repository

import (
	"add-tag/service/config"
	"add-tag/service/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AttributeBuilder(queryParameter *models.QueryParameter) (map[string]*dynamodb.AttributeValue, error) {

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	fmt.Printf("AttributeBuilder:: marshalled struct: %+v \n", av)
	return av, err
}

func MandatoryFieldsCheck(queryParameter *models.QueryParameter) error {
	if queryParameter.TagName == nil {
		err := fmt.Errorf("AttributeBuilder:: TagName is empty")
		return err
	}

	return nil
}

func TransformToQueryParamter(body *string) (*models.QueryParameter, error) {
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
	return &queryParameter, nil
}
