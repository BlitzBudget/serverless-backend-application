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
	planned := 200
	category := "category"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"planned\":" + strconv.Itoa(planned) + ",\"category\":\"" + category + "\"}"

	requestModel, err := AttributeBuilder(&body)

	assert.NoError(err)
	assert.NotNil(requestModel)

	if *requestModel.Pk != pk {
		t.Errorf("pk convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Pk, pk)
		return
	}

	if *requestModel.Sk != sk {
		t.Errorf("SK convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Sk, sk)
		return
	}

	if *requestModel.Planned != float64(planned) {
		t.Errorf("planned convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Planned, planned)
		return
	}

	if *requestModel.Category != category {
		t.Errorf("category convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Category, category)
		return
	}
}
