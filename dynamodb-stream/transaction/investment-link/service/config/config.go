package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "amount, description, category, scheduled_transaction, sk, pk, creation_date, tags"

const (
	SkPrefix = "Transaction#"
)
