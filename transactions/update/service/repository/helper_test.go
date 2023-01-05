package repository

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeBuilder(t *testing.T) {
	assert := assert.New(t)

	pk := "Wallet#2022-05-12T20:25:19Z"
	sk := "Budget#2022-05-12T20:25:19Z"
	amount := 200
	category := "category"
	description := "description"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"amount\":" + strconv.Itoa(amount) + ",\"category_id\":\"" + category + "\",\"description\":\"" + description + "\"}"

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

	if *requestModel.Amount != float64(amount) {
		t.Errorf("amount convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.Amount, amount)
		return
	}

	if *requestModel.CategoryId != category {
		t.Errorf("category convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.CategoryId, category)
		return
	}

	if *requestModel.Description != description {
		t.Errorf("description convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.Description, description)
		return
	}
}
