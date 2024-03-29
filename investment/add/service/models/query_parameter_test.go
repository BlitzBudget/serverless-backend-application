package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	pk := "Wallet#2022-05-12T14:40:29.156Z"
	sk := "Transaction#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	currentValue := 500
	investedAmount := 200
	investmentName := "investmentName"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"current_value\":" + strconv.Itoa(currentValue) + ",\"invested_amount\":" + strconv.Itoa(investedAmount) + ",\"updated_date\":\"" + updatedDate + "\",\"investment_name\":\"" + investmentName + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

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

	if *queryParameter.InvestedAmount != float64(investedAmount) {
		t.Errorf("QueryParameter: Investment Amount do not match, got = %v, want = %v \n", *queryParameter.InvestedAmount, investedAmount)
		return
	}

	if *queryParameter.InvestmentName != investmentName {
		t.Errorf("QueryParameter: Investment Name do not match, got = %v, want = %v \n", queryParameter.InvestmentName, investmentName)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v \n", err)
		return
	}
}
