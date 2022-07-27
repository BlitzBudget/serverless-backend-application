package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	amount := 20
	description := "description"
	categoryId := "Category#2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	tags := []string{"string", "string"}
	tagsString, _ := json.Marshal(tags)
	body := "{\":a\":" + strconv.Itoa(amount) + ",\":d\":\"" + description + "\",\":c\":\"" + categoryId + "\",\":u\":\"" + updatedDate + "\",\":t\":" + string(tagsString) + "}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.Amount != float64(amount) {
		t.Errorf("QueryParameter: Amount do not match, got = %v, want = %v", *queryParameter.Amount, amount)
		return
	}

	if *queryParameter.Description != description {
		t.Errorf("QueryParameter: Description do not match, got = %v, want = %v", queryParameter.Description, description)
		return
	}

	if *queryParameter.CategoryId != categoryId {
		t.Errorf("QueryParameter: Category ID do not match, got = %v, want = %v", queryParameter.CategoryId, categoryId)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: Updated Date do not match, got = %v, want = %v", queryParameter.UpdatedDate, updatedDate)
		return
	}

	if *queryParameter.Tags == nil {
		t.Errorf("QueryParameter: Tags do not match, got = %v", *queryParameter.Tags)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
