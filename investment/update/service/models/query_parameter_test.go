package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	investName := "investmentName"
	investedAmount := 20
	currentValue := 10
	updatedDate := "2022-05-12T20:25:19Z"
	body := "{\":n\":\"" + investName + "\",\":a\":" + strconv.Itoa(investedAmount) + ",\":c\":" + strconv.Itoa(currentValue) + ",\":u\":\"" + updatedDate + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.InvestmentName != investName {
		t.Errorf("QueryParameter: Investment Name do not match, got = %v, want = %v", *queryParameter.InvestmentName, investName)
		return
	}

	if queryParameter.InvestedAmount != sk {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v", queryParameter.Sk, sk)
		return
	}

	if *queryParameter.CreationDate != creationDate {
		t.Errorf("QueryParameter: Creation Date do not match, got = %v, want = %v", queryParameter.CreationDate, creationDate)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: Updated Date do not match, got = %v, want = %v", queryParameter.UpdatedDate, updatedDate)
		return
	}

	if *queryParameter.Category != categoryId {
		t.Errorf("QueryParameter: Category ID do not match, got = %v, want = %v", *queryParameter.Category, categoryId)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
