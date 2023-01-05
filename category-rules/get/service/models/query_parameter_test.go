package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	wallet := "Category#2022-04-12T14:40:29.156Z"
	category := "Wallet#2022-05-12T14:40:29.156Z"
	body := "{\"wallet_id\":\"" + wallet + "\",\"category_id\":\"" + category + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.WalletId != wallet {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v \n", *queryParameter.WalletId, wallet)
		return
	}

	if *queryParameter.CategoryId != category {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v \n", queryParameter.CategoryId, category)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v \n", err)
		return
	}
}
