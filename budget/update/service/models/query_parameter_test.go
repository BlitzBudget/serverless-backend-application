package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	planned := 2000
	category := "category"
	updatedDate := "2022-05-12T20:25:19Z"
	body := "{\":p\":" + strconv.Itoa(planned) + ",\":c\":\"" + category + "\",\":u\":\"" + updatedDate + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.Planned != float64(planned) {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v", *queryParameter.Planned, planned)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: UpdatedDate do not match, got = %v, want = %v", queryParameter.UpdatedDate, updatedDate)
		return
	}

	if *queryParameter.Category != category {
		t.Errorf("QueryParameter: Category do not match, got = %v, want = %v", queryParameter.Category, category)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
