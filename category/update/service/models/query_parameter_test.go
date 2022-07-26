package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	categoryName := "Category#2022-04-12T14:40:29.156Z"
	categoryType := "Expense"
	updatedDate := "2022-05-12T20:25:19Z"
	body := "{\":n\":\"" + categoryName + "\",\":t\":\"" + categoryType + "\",\":u\":\"" + updatedDate + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.CategoryName != categoryName {
		t.Errorf("QueryParameter: CategoryName do not match, got = %v, want = %v", *queryParameter.CategoryName, categoryName)
		return
	}

	if *queryParameter.CategoryType != categoryType {
		t.Errorf("QueryParameter: CategoryType do not match, got = %v, want = %v", queryParameter.CategoryType, categoryType)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: Updated Date do not match, got = %v, want = %v", queryParameter.UpdatedDate, updatedDate)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
