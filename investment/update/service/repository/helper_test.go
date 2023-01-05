package repository

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeBuilder(t *testing.T) {
	assert := assert.New(t)

	pk := "pk"
	sk := "sk"
	investedAmount := 1000
	investmentName := "investmentName"
	currentAmount := 2000
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"invested_amount\":" + strconv.Itoa(investedAmount) + ",\"investment_name\":\"" + investmentName + "\",\"current_value\":" + strconv.Itoa(currentAmount) + "}"

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

	if *requestModel.InvestedAmount != float64(investedAmount) {
		t.Errorf("InvestedAmount convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.InvestedAmount, investedAmount)
		return
	}

	if *requestModel.CurrentValue != float64(currentAmount) {
		t.Errorf("CurrentValue convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.CurrentValue, currentAmount)
		return
	}

	if *requestModel.InvestmentName != investmentName {
		t.Errorf("InvestmentName convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.InvestmentName, investmentName)
		return
	}
}
