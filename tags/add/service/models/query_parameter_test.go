package models

import (
	"encoding/json"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	pk := "Wallet#2022-05-12T14:40:29.156Z"
	sk := "Tag#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	tagName := "tagName"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"updated_date\":\"" + updatedDate + "\",\"tag_name\":\"" + tagName + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.Pk != pk {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v", *queryParameter.Pk, pk)
		return
	}

	if queryParameter.Sk != sk {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v", queryParameter.Sk, sk)
		return
	}

	if *queryParameter.CreationDate != creationDate {
		t.Errorf("QueryParameter: Creation Date do not match, got = %v, want = %v", queryParameter.CreationDate, creationDate)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: Updated Date do not match, got = %v, want = %v", queryParameter.UpdatedDate, updatedDate)
		return
	}

	if *queryParameter.TagName != tagName {
		t.Errorf("QueryParameter: Tag Name do not match, got = %v, want = %v", queryParameter.TagName, tagName)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
