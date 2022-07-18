package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	pk := "Wallet#2022-05-12T14:40:29.156Z"
	sk := "Transaction#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	amount := 100
	description := "description"
	categoryId := "categoryId"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"amount\":\"" + strconv.Itoa(amount) + "\",\"description\":\"" + description + "\",\"updated_date\":\"" + updatedDate + "\",\"category_id\":\"" + categoryId + "\"}"
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

	if *queryParameter.Amount != float64(amount) {
		t.Errorf("QueryParameter: Amount do not match, got = %v, want = %v", *queryParameter.Amount, amount)
		return
	}

	if *queryParameter.CategoryId != categoryId {
		t.Errorf("QueryParameter: Category ID do not match, got = %v, want = %v", *queryParameter.CategoryId, categoryId)
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
