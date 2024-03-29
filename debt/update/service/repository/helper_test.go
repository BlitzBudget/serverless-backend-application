package repository

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeBuilder(t *testing.T) {
	assert :=  assert.New(t)

	pk := "pk"
	sk := "sk"
	debtName := "debtName"
	debtRepaid := true
	debtedAmount := 1000
	currentAmount := 2000
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"debt_name\":\"" + debtName + "\",\"debt_repaid\":" + strconv.FormatBool(debtRepaid) + ",\"debted_amount\":" + strconv.Itoa(debtedAmount) + ",\"current_value\":" + strconv.Itoa(currentAmount) + "}"

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

	if *requestModel.DebtName != debtName {
		t.Errorf("CategoryName convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.DebtName, debtName)
		return
	}

	if *requestModel.DebtRepaid != debtRepaid {
		t.Errorf("DebtRepaid convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.DebtRepaid, debtRepaid)
		return
	}
}
