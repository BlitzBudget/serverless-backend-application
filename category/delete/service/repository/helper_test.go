package repository

import (
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	pk := "Wallet#2022-05-12T20:25:19Z"
	sk := "Budget#2022-05-12T20:25:19Z"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\"}"

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

}
