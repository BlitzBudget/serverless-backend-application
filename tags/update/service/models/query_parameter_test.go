package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	tagName := "tagName"
	updatedDate := "2022-05-12T20:25:19Z"
	body := "{\":n\":\"" + tagName + "\",\":u\":\"" + updatedDate + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.TagName != tagName {
		t.Errorf("QueryParameter: Tag Name do not match, got = %v, want = %v", *queryParameter.TagName, tagName)
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
