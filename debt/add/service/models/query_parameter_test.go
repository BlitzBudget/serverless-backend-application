package models

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryParameter(t *testing.T) {
	assert := assert.New(t)

	pk := "Wallet#2022-05-12T14:40:29.156Z"
	sk := "Transaction#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	currentValue := 2000
	debtedAmount := 1000
	debtName := "debtName"
	debtRepaid := true
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"current_value\":" + strconv.Itoa(currentValue) + ",\"debted_amount\":" + strconv.Itoa(debtedAmount) + ",\"debt_name\": \"" + debtName + "\",\"updated_date\":\"" + updatedDate + "\",\"debt_repaid\":" + strconv.FormatBool(debtRepaid) + "}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	assert.NoError(err)

	if *queryParameter.Pk != pk {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v \n", *queryParameter.Pk, pk)
		return
	}

	if queryParameter.Sk != sk {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v \n", queryParameter.Sk, sk)
		return
	}

	if *queryParameter.CreationDate != creationDate {
		t.Errorf("QueryParameter: Creation Date do not match, got = %v, want = %v \n", queryParameter.CreationDate, creationDate)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: Updated Date do not match, got = %v, want = %v \n", queryParameter.UpdatedDate, updatedDate)
		return
	}

	if *queryParameter.CurrentValue != float64(currentValue) {
		t.Errorf("QueryParameter: CurrentValue do not match, got = %v, want = %v \n", *queryParameter.CurrentValue, currentValue)
		return
	}

	if *queryParameter.DebtedAmount != float64(debtedAmount) {
		t.Errorf("QueryParameter: DebtedAmount do not match, got = %v, want = %v \n", *queryParameter.DebtedAmount, debtedAmount)
		return
	}

	if *queryParameter.DebtName != debtName {
		t.Errorf("QueryParameter: Debt Name do not match, got = %v, want = %v \n", queryParameter.DebtName, debtName)
		return
	}

	if *queryParameter.DebtRepaid != debtRepaid {
		t.Errorf("QueryParameter: Debt Repaid do not match, got = %v, want = %v \n", queryParameter.DebtRepaid, debtRepaid)
		return
	}
}
