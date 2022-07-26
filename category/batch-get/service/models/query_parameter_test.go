package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	userId := "Category#2022-04-12T14:40:29.156Z"
	categoryId := "Wallet#2022-05-12T14:40:29.156Z"
	categoryIds := []string{categoryId}
	categoryIdString, _ := json.Marshal(categoryIds)
	body := "{\"user_id\":\"" + userId + "\",\"category_ids\":" + string(categoryIdString) + "}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.UserId != userId {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v", *queryParameter.UserId, userId)
		return
	}

	if len(*queryParameter.CategoryIds) != 1 {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v", len(*queryParameter.CategoryIds), 1)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
