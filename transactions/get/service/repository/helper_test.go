package repository

import (
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	endsWithDate := "2022-05-12T20:25:19Z"
	startsWithDate := "2022-05-12T20:25:19Z"
	walletId := "walletId"
	body := "{\"wallet_id\":\"" + walletId + "\",\"starts_with_date\":\"" + startsWithDate + "\",\"ends_with_date\":\"" + endsWithDate + "\"}"

	requestModel, err := AttributeBuilder(&body)

	if err != nil {
		t.Errorf("Err is null")
		return
	}

	if requestModel == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *requestModel.EndsWithDate != endsWithDate {
		t.Errorf("EndsWithDate convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.EndsWithDate, endsWithDate)
		return
	}

	if *requestModel.StartsWithDate != startsWithDate {
		t.Errorf("StartsWithDate convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.StartsWithDate, startsWithDate)
		return
	}

	if *requestModel.WalletId != walletId {
		t.Errorf("WalletId convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.WalletId, walletId)
		return
	}
}
