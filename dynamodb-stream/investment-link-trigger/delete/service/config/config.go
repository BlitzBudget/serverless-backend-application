package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "debt_id, transaction_amount, transaction_name, sk, pk, creation_date"

const (
	SkDebtRulePrefix = "DebtRule#"
)
