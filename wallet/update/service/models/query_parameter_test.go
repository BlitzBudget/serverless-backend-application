package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	currency := "$"
	name := "name"
	updatedate := "2022-10-10"
	body := "{\":c\":\"" + currency + "\",\":n\":\"" + name + "\",\":u\":\"" + updatedate + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.Currency != currency {
		t.Errorf("QueryParameter: Currency do not match, got = %v, want = %v \n", *queryParameter.Currency, currency)
		return
	}

	if *queryParameter.Name != name {
		t.Errorf("QueryParameter: Name do not match, got = %v, want = %v \n", queryParameter.Name, name)
		return
	}

	if *queryParameter.UpdatedDate != updatedate {
		t.Errorf("QueryParameter: Updated Date do not match, got = %v, want = %v \n", queryParameter.UpdatedDate, updatedate)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v \n", err)
		return
	}
}
