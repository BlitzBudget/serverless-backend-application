package repository

import (
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	walletId := "walletId"
	body := "{\"wallet_id\":\"" + walletId + "\"}"

	requestModel, err := AttributeBuilder(&body)

	if err != nil {
		t.Errorf("Err is null")
		return
	}

	if requestModel == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *requestModel.WalletId != walletId {
		t.Errorf("WalletId convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.WalletId, walletId)
		return
	}
}
