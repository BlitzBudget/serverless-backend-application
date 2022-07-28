package models

import (
	"encoding/json"
	"testing"
)

func TestConfig(t *testing.T) {
	pk := "Wallet#2022-05-12T14:40:29.156Z"
	sk := "Transaction#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	category := "category"
	tags := []string{"tags", "tag"}
	description := "description"
	tagsByte, _ := json.Marshal(tags)
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"category\":\"" + category + "\",\"description\":\"" + description + "\",\"tags\": " + string(tagsByte) + "}"
	queryParameter := Transaction{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.Pk != pk {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v", *queryParameter.Pk, pk)
		return
	}

	if *queryParameter.Sk != sk {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v", queryParameter.Sk, sk)
		return
	}

	if *queryParameter.CreationDate != creationDate {
		t.Errorf("QueryParameter: Creation Date do not match, got = %v, want = %v", queryParameter.CreationDate, creationDate)
		return
	}

	if *queryParameter.Category != category {
		t.Errorf("QueryParameter: Category do not match, got = %v, want = %v", queryParameter.Category, category)
		return
	}

	if *queryParameter.Tags == nil {
		t.Errorf("QueryParameter: Tags do not match, got = %v", *queryParameter.Tags)
		return
	}

	if *queryParameter.Description != description {
		t.Errorf("QueryParameter: Description do not match, got = %v, want = %v", *queryParameter.Description, description)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}

}
