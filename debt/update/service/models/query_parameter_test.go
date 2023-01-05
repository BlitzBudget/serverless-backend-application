package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	debtedAmount := 30
	debtName := "debtName"
	currentValue := 10
	updatedDate := "2022-05-12T20:25:19Z"
	debtRepaid := true
	body := "{\":a\":" + strconv.Itoa(debtedAmount) + ",\":n\":\"" + debtName + "\",\":c\":" + strconv.Itoa(currentValue) + ",\":u\":\"" + updatedDate + "\",\":r\":" + strconv.FormatBool(debtRepaid) + "}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.DebtedAmount != float64(debtedAmount) {
		t.Errorf("QueryParameter: Debted Amount do not match, got = %v, want = %v \n", *queryParameter.DebtedAmount, debtedAmount)
		return
	}

	if *queryParameter.DebtName != debtName {
		t.Errorf("QueryParameter: DebtName do not match, got = %v, want = %v \n", queryParameter.DebtName, debtName)
		return
	}

	if *queryParameter.CurrentValue != float64(currentValue) {
		t.Errorf("QueryParameter: Current Value do not match, got = %v, want = %v \n", queryParameter.CurrentValue, currentValue)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: Updated Date do not match, got = %v, want = %v \n", queryParameter.UpdatedDate, updatedDate)
		return
	}

	if *queryParameter.DebtRepaid != debtRepaid {
		t.Errorf("QueryParameter: Debt Repaid do not match, got = %v, want = %v \n", *queryParameter.DebtRepaid, debtRepaid)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v \n", err)
		return
	}
}
