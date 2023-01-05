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
	transactionName := "transactionName"
	debtId := "debtId"
	transactionAmount := 2000
	body := "{\"pk\":\"" + pk + "\",\"creation_date\":\"" + creationDate + "\",\"transaction_name\":\"" + transactionName + "\",\"debt_id\":\"" + debtId + "\",\"updated_date\":\"" + updatedDate + "\",\"transaction_amount\":" + strconv.Itoa(transactionAmount) + "}"

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

	if *(*got["transaction_name"]).S != transactionName {
		t.Errorf("name transaction_name to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["transaction_name"]).S, transactionName)
		return
	}

	if *(*got["debt_id"]).S != debtId {
		t.Errorf("name debt_id to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["debt_id"]).S, debtId)
		return
	}

	if *(*got["transaction_amount"]).N != strconv.Itoa(transactionAmount) {
		t.Errorf("name transaction_amount to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["transaction_amount"]).N, transactionAmount)
		return
	}
}
