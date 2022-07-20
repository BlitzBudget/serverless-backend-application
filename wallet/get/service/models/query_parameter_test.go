package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	userId := "userId"
	body := "{\"user_id\":\"" + userId + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.UserId != userId {
		t.Errorf("UserId convertion to DynamoDB attribute not correct, got = %v, want = %v", *queryParameter.UserId, userId)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
