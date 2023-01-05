package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	wallet := "Wallet#2022-04-12T14:40:29.156Z"
	body := "{\"wallet_id\":\"" + wallet + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.WalletId != wallet {
		t.Errorf("QueryParameter: WalletId do not match, got = %v, want = %v", *queryParameter.WalletId, wallet)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
