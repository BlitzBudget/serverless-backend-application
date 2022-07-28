package repository

import (
	"errors"
	"testing"
	"update-category/service/config"
	"update-category/service/models"

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
	categoryName := "categoryName"
	categoryType := "Expense"
	updateDate := "2022-10-09"
	av, err := dynamodbattribute.MarshalMap(models.QueryParameter{
		CategoryName: &categoryName,
		CategoryType: &categoryType,
		UpdatedDate:  &updateDate,
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
