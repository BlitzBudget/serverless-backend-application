package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeBuilder(t *testing.T) {
	assert := assert.New(t)

	pk := "Wallet#2022-05-12T20:25:19Z"
	sk := "Budget#2022-05-12T20:25:19Z"
	currency := "$"
	name := "name"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"currency\":\"" + currency + "\",\"name\":\"" + name + "\"}"

	requestModel, err := AttributeBuilder(&body)
	assert.NoError(err)
	assert.NotNil(requestModel)

	if *requestModel.Pk != pk {
		t.Errorf("pk convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.Pk, pk)
		return
	}

	if *requestModel.Sk != sk {
		t.Errorf("SK convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.Sk, sk)
		return
	}

	if *requestModel.Currency != currency {
		t.Errorf("Currency convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.Currency, currency)
		return
	}

	if *requestModel.Name != name {
		t.Errorf("Name convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.Name, name)
		return
	}
}
