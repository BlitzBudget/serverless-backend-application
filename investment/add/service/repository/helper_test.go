package repository

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeBuilder(t *testing.T) {
	assert := assert.New(t)

	pk := "Wallet#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	currentValue := 2000
	investedAmount := 1000
	investmentName := "investmentName"
	body := "{\"pk\":\"" + pk + "\",\"creation_date\":\"" + creationDate + "\",\"current_value\":" + strconv.Itoa(currentValue) + ",\"invested_amount\":" + strconv.Itoa(investedAmount) + ",\"updated_date\":\"" + updatedDate + "\",\"investment_name\":\"" + investmentName + "\"}"

	got, err := AttributeBuilder(&body)
	assert.NoError(err)
	assert.NotNil(got)

	if *(*got["creation_date"]).S == "" {
		t.Errorf("name creationDate to DynamoDB attribute not correct, got = %v \n", *(*got["creation_date"]).S)
		return
	}

	if *(*got["updated_date"]).S == "" {
		t.Errorf("name updated_date to DynamoDB attribute not correct, got = %v \n", *(*got["updated_date"]).S)
		return
	}

	if *(*got["pk"]).S != pk {
		t.Errorf("name pk to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["pk"]).S, pk)
		return
	}

	if *(*got["sk"]).S == "" {
		t.Errorf("name sk to DynamoDB attribute not correct, got = %v \n", *(*got["sk"]).S)
		return
	}

	if *(*got["current_value"]).N != strconv.Itoa(currentValue) {
		t.Errorf("name current_value to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["current_value"]).N, currentValue)
		return
	}

	if *(*got["invested_amount"]).N != strconv.Itoa(investedAmount) {
		t.Errorf("name invested_amount to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["invested_amount"]).N, investedAmount)
		return
	}

	if *(*got["investment_name"]).S != investmentName {
		t.Errorf("name investment_name to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["investment_name"]).S, investmentName)
		return
	}

}
