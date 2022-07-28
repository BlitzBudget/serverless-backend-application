package repository

import (
	"add-debt-link/service/config"
	"add-debt-link/service/models"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (m *mockDynamodbClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	result := dynamodb.GetItemOutput{}

	if input.TableName == nil || *input.TableName == "" {
		return &result, errors.New("You must supply a table name")
	}

	pk := "pk"
	sk := "sk"
	debtName := "debtName"
	debtRepaid := false
	creationDate := "creationDate"
	debtedAmount := float64(20)
	currentValue := float64(30)
	item := models.Debt{
		Pk:           &pk,
		Sk:           &sk,
		CreationDate: &creationDate,
		DebtedAmount: &debtedAmount,
		CurrentValue: &currentValue,
		DebtName:     &debtName,
		DebtRepaid:   &debtRepaid,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return &result, err
	}

	result = dynamodb.GetItemOutput{
		Item: av,
	}

	return &result, nil
}

func TestGetItem(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	t.Log("Starting unit test at " + nowString)

	config.TableName = "tablename"
	pk := "pk"
	sk := "sk"
	mockSvc := &mockDynamodbClient{}

	item, err := GetDebtItem(mockSvc, &pk, &sk)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Retrieved movie '" + *item.Pk + "' from table ")
}
