package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestInvestmentRule(t *testing.T) {
	pk := "Wallet#2022-05-12T14:40:29.156Z"
	sk := "Investment#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	investmentId := "20"
	transactionAmount := 10
	transactionName := "transactionName"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"investment_id\":\"" + investmentId + "\",\"transaction_amount\":" + strconv.Itoa(transactionAmount) + ",\"transaction_name\":\"" + transactionName + "\"}"
	queryParameter := InvestmentRule{}
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

	if *queryParameter.TransactionAmount != float64(transactionAmount) {
		t.Errorf("QueryParameter: transaction amount do not match, got = %v, want = %v", queryParameter.TransactionAmount, transactionAmount)
		return
	}

	if *queryParameter.InvestmentId != investmentId {
		t.Errorf("QueryParameter: Investment ID do not match, got = %v, want = %v", queryParameter.InvestmentId, investmentId)
		return
	}

	if *queryParameter.TransactionName != transactionName {
		t.Errorf("QueryParameter: Transaction Name do not match, got = %v, want = %v", *queryParameter.TransactionName, transactionName)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
