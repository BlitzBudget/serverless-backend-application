package repository

import (
	"add-investment-link/service/config"
	"add-investment-link/service/models"
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
	creationDate := "creationDate"
	investAmount := float64(20)
	currentValue := float64(30)
	investmentName := "investmentName"
	item := models.Investment{
		Pk:             &pk,
		Sk:             &sk,
		CreationDate:   &creationDate,
		InvestedAmount: &investAmount,
		InvestmentName: &investmentName,
		CurrentValue:   &currentValue,
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

	item, err := GetInvestmentItem(mockSvc, &pk, &sk)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Retrieved movie '" + *item.Pk + "' from table ")
}

func TestGetRuleItem(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	t.Log("Starting unit test at " + nowString)

	config.TableName = "tablename"
	pk := "pk"
	sk := "sk"
	mockSvc := &mockDynamodbClient{}

	item, err := GetInvestmentRuleItem(mockSvc, &pk, &sk)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Retrieved movie '" + *item.Pk + "' from table ")
}
