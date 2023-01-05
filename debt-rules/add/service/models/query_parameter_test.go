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
	transactionName := "TransactionName"
	debtId := "debtId"
	transactionAmount := 1000
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"transaction_name\":\"" + transactionName + "\",\"debt_id\":\"" + debtId + "\",\"transaction_amount\": " + strconv.Itoa(transactionAmount) + ",\"updated_date\":\"" + updatedDate + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.Pk != pk {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v \n", *queryParameter.Pk, pk)
		return
	}

	if queryParameter.Sk != sk {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v \n", queryParameter.Sk, sk)
		return
	}

	if *queryParameter.CreationDate != creationDate {
		t.Errorf("QueryParameter: Creation Date do not match, got = %v, want = %v \n", queryParameter.CreationDate, creationDate)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: Updated Date do not match, got = %v, want = %v \n", queryParameter.UpdatedDate, updatedDate)
		return
	}

	if *queryParameter.TransactionName != transactionName {
		t.Errorf("QueryParameter: TransactionName do not match, got = %v, want = %v \n", *queryParameter.TransactionName, transactionName)
		return
	}

	if *queryParameter.DebtId != debtId {
		t.Errorf("QueryParameter: Debt Id do not match, got = %v, want = %v \n", *queryParameter.DebtId, debtId)
		return
	}

	if *queryParameter.TransactionAmount != float64(transactionAmount) {
		t.Errorf("QueryParameter: Transaction Amount do not match, got = %v, want = %v \n", queryParameter.TransactionAmount, transactionAmount)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v \n", err)
		return
	}
}
