package repository

import (
	"strconv"
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	pk := "Wallet#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	transactionName := "transactionName"
	goalId := "goalId"
	transactionAmount := 2000
	body := "{\"pk\":\"" + pk + "\",\"creation_date\":\"" + creationDate + "\",\"transaction_name\":\"" + transactionName + "\",\"transaction_amount\":" + strconv.Itoa(transactionAmount) + ",\"updated_date\":\"" + updatedDate + "\",\"goal_id\":\"" + goalId + "\"}"

	got, err := AttributeBuilder(&body)
	if err != nil {
		t.Errorf("AttributeBuilder() error = %v", err)
		return
	}

	if got == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *(*got["creation_date"]).S == "" {
		t.Errorf("name creationDate to DynamoDB attribute not correct, got = %v", *(*got["creation_date"]).S)
		return
	}

	if *(*got["updated_date"]).S == "" {
		t.Errorf("name updated_date to DynamoDB attribute not correct, got = %v", *(*got["updated_date"]).S)
		return
	}

	if *(*got["pk"]).S != pk {
		t.Errorf("name pk to DynamoDB attribute not correct, got = %v, want = %v", *(*got["pk"]).S, pk)
		return
	}

	if *(*got["sk"]).S == "" {
		t.Errorf("name sk to DynamoDB attribute not correct, got = %v", *(*got["sk"]).S)
		return
	}

	if *(*got["transaction_name"]).S != transactionName {
		t.Errorf("name transaction_name to DynamoDB attribute not correct, got = %v, want = %v", *(*got["transaction_name"]).S, transactionName)
		return
	}

	if *(*got["goal_id"]).S != goalId {
		t.Errorf("name goal_id to DynamoDB attribute not correct, got = %v, want = %v", *(*got["goal_id"]).S, goalId)
		return
	}

	if *(*got["transaction_amount"]).N != strconv.Itoa(transactionAmount) {
		t.Errorf("name transaction_amount to DynamoDB attribute not correct, got = %v, want = %v", *(*got["transaction_amount"]).N, transactionAmount)
		return
	}
}
