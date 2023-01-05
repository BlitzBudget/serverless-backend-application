package repository

import (
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	userId := "userId"
	body := "{\"user_id\":\"" + userId + "\"}"

	requestModel, err := AttributeBuilder(&body)

	if err != nil {
		t.Errorf("Err is null")
		return
	}

	if requestModel == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *requestModel.UserId != userId {
		t.Errorf("UserId convertion to DynamoDB attribute not correct, got = %v, want = %v \n", *requestModel.UserId, userId)
		return
	}
}
