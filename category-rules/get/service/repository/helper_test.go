package repository

import (
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	categoryId := "categoryId"
	walletId := "walletId"
	body := "{\"wallet_id\":\"" + walletId + "\",\"category_id\":\"" + categoryId + "\"}"

	requestModel, err := AttributeBuilder(&body)

	if err != nil {
		t.Errorf("Err is null")
		return
	}

	if requestModel == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *requestModel.CategoryId != categoryId {
		t.Errorf("CategoryId convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.CategoryId, categoryId)
		return
	}

	if *requestModel.WalletId != walletId {
		t.Errorf("WalletId convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.WalletId, walletId)
		return
	}
}
