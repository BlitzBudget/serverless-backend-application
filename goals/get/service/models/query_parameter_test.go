package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	walletId := "Wallet#2022-04-12T14:40:29.156Z"

	body := "{\"wallet_id\":\"" + walletId + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.WalletId != walletId {
		t.Errorf("QueryParameter: walletId do not match, got = %v, want = %v \n", *queryParameter.WalletId, walletId)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v \n", err)
		return
	}
}
