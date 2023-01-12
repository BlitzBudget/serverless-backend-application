package repository

import (
	"overview/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
)

type mockDynamodbClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamodbClient) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	result := dynamodb.QueryOutput{}

	wallet := "wallet"
	startsWithDate := "2022-10-20"
	endsWithDate := "2022-12-01"
	item := models.QueryParameter{
		WalletId:       &wallet,
		StartsWithDate: &startsWithDate,
		EndsWithDate:   &endsWithDate,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return &result, err
	}

	result = dynamodb.QueryOutput{
		Items: []map[string]*dynamodb.AttributeValue{
			av,
		},
	}

	return &result, nil
}

func TestQueryItem(t *testing.T) {
	mockSvc := &mockDynamodbClient{}
	wallet := "wallet"
	startsWithDate := "2022-10-20"
	endsWithDate := "2022-12-01"
	av := models.QueryParameter{
		WalletId:       &wallet,
		StartsWithDate: &startsWithDate,
		EndsWithDate:   &endsWithDate,
	}

	items, err := QueryItem(&av, mockSvc)

	assert := assert.New(t)
	assert.NoError(err)
	assert.NotNil(items)

}
