package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	wallet := "Wallet#2022-05-12T14:40:29.156Z"
	startsWithDate := "2022-05-12T20:25:19Z"
	endsWithDate := "2022-05-12T20:25:19Z"
	body := "{\"wallet_id\":\"" + wallet + "\",\"starts_with_date\":\"" + startsWithDate + "\",\"ends_with_date\":\"" + endsWithDate + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.WalletId != wallet {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v \n", *queryParameter.WalletId, wallet)
		return
	}

	if *queryParameter.EndsWithDate != endsWithDate {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v \n", queryParameter.EndsWithDate, endsWithDate)
		return
	}

	if *queryParameter.StartsWithDate != startsWithDate {
		t.Errorf("QueryParameter: Creation Date do not match, got = %v, want = %v \n", queryParameter.StartsWithDate, startsWithDate)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v \n", err)
		return
	}
}
