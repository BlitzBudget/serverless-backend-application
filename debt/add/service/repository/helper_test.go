package repository

import (
	"strconv"
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	pk := "Wallet#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	debtName := "debtName"
	debtRepaid := true
	debtedAmount := 2000
	currentValue := 1000
	body := "{\"pk\":\"" + pk + "\",\"creation_date\":\"" + creationDate + "\",\"debt_name\":\"" + debtName + "\",\"debt_repaid\":" + strconv.FormatBool(debtRepaid) + ",\"updated_date\":\"" + updatedDate + "\",\"debted_amount\":" + strconv.Itoa(debtedAmount) + ",\"current_value\":" + strconv.Itoa(currentValue) + "}"

	got, err := AttributeBuilder(&body)
	if err != nil {
		t.Errorf("AttributeBuilder() error = %v \n", err)
		return
	}

	if got == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

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

	if *(*got["debt_name"]).S != debtName {
		t.Errorf("name debt_name to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["debt_name"]).S, debtName)
		return
	}

	if *(*got["debt_repaid"]).BOOL != debtRepaid {
		t.Errorf("name debt_repaid to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["debt_repaid"]).S, debtRepaid)
		return
	}

	if *(*got["current_value"]).N != strconv.Itoa(currentValue) {
		t.Errorf("name current_value to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["current_value"]).N, currentValue)
		return
	}

	if *(*got["debted_amount"]).N != strconv.Itoa(debtedAmount) {
		t.Errorf("name debted_amount to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["debted_amount"]).N, debtedAmount)
		return
	}
}
