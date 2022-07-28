package repository

import (
	"errors"
	"patch-budget/service/config"
	"patch-budget/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Define a mock struct to use in unit tests
type mockDynamodbClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamodbClient) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	result := dynamodb.UpdateItemOutput{}

	if input.TableName == nil || *input.TableName == "" {
		return &result, errors.New("Missing required field UpdateItemInput.TableName")
	}

	return &result, nil
}

func TestUpdateItem(t *testing.T) {
	planned := float64(20)
	category := "category"
	updateDate := "2022-10-09"
	av, err := dynamodbattribute.MarshalMap(models.QueryParameter{
		Planned:     &planned,
		Category:    &category,
		UpdatedDate: &updateDate,
	})
	if err != nil {
		t.Errorf("UpdateItem() error = %v", err)
	}

	pk := "pk"
	sk := "sk"
	requestModel := models.RequestModel{
		Pk: &pk,
		Sk: &sk,
	}

	mockSvc := &mockDynamodbClient{}

	config.TableName = "tablename"

	UpdateItem(av, mockSvc, &requestModel)
}
