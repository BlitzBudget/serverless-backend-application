package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	wallet := "Wallet#2022-04-12T14:40:29.156Z"
	startsWithDate := "2022-05-12T14:40:29.156Z"
	endsWithDate := "2022-05-12T20:25:19Z"
	body := "{\"wallet_id\":\"" + wallet + "\",\"starts_with_date\":\"" + startsWithDate + "\",\"ends_with_date\":\"" + endsWithDate + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.WalletId != wallet {
		t.Errorf("QueryParameter: Wallet do not match, got = %v, want = %v", *queryParameter.WalletId, wallet)
		return
	}

	if *queryParameter.StartsWithDate != startsWithDate {
		t.Errorf("QueryParameter: StartsWithDate do not match, got = %v, want = %v", queryParameter.StartsWithDate, startsWithDate)
		return
	}

	if *queryParameter.EndsWithDate != endsWithDate {
		t.Errorf("QueryParameter: Ends With Date do not match, got = %v, want = %v", queryParameter.EndsWithDate, endsWithDate)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
