package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	userId := "User#2022-04-12T14:40:29.156Z"
	body := "{\"user_id\":\"" + userId + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.UserId != userId {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v", *queryParameter.UserId, userId)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
