package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestInvestment(t *testing.T) {
	pk := "Wallet#2022-05-12T14:40:29.156Z"
	sk := "Investment#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	investedAmount := 20
	currentValue := 10
	investmentName := "investmentName"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"invested_amount\":" + strconv.Itoa(investedAmount) + ",\"current_value\":" + strconv.Itoa(currentValue) + ",\"investment_name\":\"" + investmentName + "\"}"
	queryParameter := Investment{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.Pk != pk {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v", *queryParameter.Pk, pk)
		return
	}

	if *queryParameter.Sk != sk {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v", queryParameter.Sk, sk)
		return
	}

	if *queryParameter.CreationDate != creationDate {
		t.Errorf("QueryParameter: Creation Date do not match, got = %v, want = %v", queryParameter.CreationDate, creationDate)
		return
	}

	if *queryParameter.InvestedAmount != float64(investedAmount) {
		t.Errorf("QueryParameter: Invested Amount do not match, got = %v, want = %v", queryParameter.InvestedAmount, investedAmount)
		return
	}

	
	if *queryParameter.CurrentValue != float64(currentValue) {
		t.Errorf("QueryParameter: Current Value do not match, got = %v, want = %v", queryParameter.CurrentValue, currentValue)
		return
	}

	if *queryParameter.InvestmentName != investmentName {
		t.Errorf("QueryParameter: Investment Name do not match, got = %v, want = %v", *queryParameter.InvestmentName, investmentName)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
