package repository

import (
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	pk := "Wallet#2022-05-12T20:25:19Z"
	sk := "Budget#2022-05-12T20:25:19Z"
	currency := "$"
	name := "name"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"currency\":\"" + currency + "\",\"name\":\"" + name + "\"}"

	requestModel := AttributeBuilder(&body)

	if requestModel == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *requestModel.Pk != pk {
		t.Errorf("pk convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Pk, pk)
		return
	}

	if *requestModel.Sk != sk {
		t.Errorf("SK convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Sk, sk)
		return
	}

	if *requestModel.Currency != currency {
		t.Errorf("Currency convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Currency, currency)
		return
	}

	if *requestModel.Name != name {
		t.Errorf("Name convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Name, name)
		return
	}
}
