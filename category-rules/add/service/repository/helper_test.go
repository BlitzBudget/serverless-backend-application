package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeBuilder(t *testing.T) {
	assert := assert.New(t)

	pk := "Wallet#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	transactionName := "transactionName"
	categoryId := "categoryId"
	body := "{\"pk\":\"" + pk + "\",\"creation_date\":\"" + creationDate + "\",\"category_id\":\"" + categoryId + "\",\"transaction_name\":\"" + transactionName + "\",\"updated_date\":\"" + updatedDate + "\"}"

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

	if *(*got["category_id"]).S != categoryId {
		t.Errorf("name category_id to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["category_id"]).S, categoryId)
		return
	}
}
