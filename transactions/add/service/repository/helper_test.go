package repository

import (
	"strconv"
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	pk := "Wallet#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	amount := 100
	description := "description"
	categoryId := "categoryId"
	body := "{\"pk\":\"" + pk + "\",\"creation_date\":\"" + creationDate + "\",\"amount\":" + strconv.Itoa(amount) + ",\"description\":\"" + description + "\",\"category_id\":\"" + categoryId + "\",\"updated_date\":\"" + updatedDate + "\"}"

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

	if *(*got["amount"]).N != strconv.Itoa(amount) {
		t.Errorf("name amount to DynamoDB attribute not correct, got = %v, want = %v", *(*got["amount"]).N, amount)
		return
	}

	if *(*got["description"]).S != description {
		t.Errorf("name description to DynamoDB attribute not correct, got = %v, want = %v", *(*got["description"]).S, description)
		return
	}

	if *(*got["category_id"]).S != categoryId {
		t.Errorf("name category_id to DynamoDB attribute not correct, got = %v, want = %v", *(*got["category_id"]).S, categoryId)
		return
	}
}
