package repository

import (
	"add-debt-link/service/config"
	"add-debt-link/service/models"
	"errors"
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
	amount := float64(20)
	category := "category"
	creationDate := "2022-10-09"
	av, err := dynamodbattribute.MarshalMap(models.Transaction{
		Amount:       &amount,
		Category:     &category,
		CreationDate: &creationDate,
	})
	if err != nil {
		t.Errorf("UpdateItem() error = %v", err)
	}

	pk := "pk"
	sk := "sk"
	updateExpression := "updateExpression"

	mockSvc := &mockDynamodbClient{}

	config.TableName = "tablename"

	UpdateItem(av, mockSvc, &pk, &sk, updateExpression)
}
