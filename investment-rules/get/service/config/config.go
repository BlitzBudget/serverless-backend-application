package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "transaction_amount, transaction_name, sk, pk, creation_date"

const (
	SkPrefix = "InvestmentRule#"
)
