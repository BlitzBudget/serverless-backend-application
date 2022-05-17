package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "debt_amount, current_value, debt_name, sk, pk, creation_date"

const (
	SkPrefix = "Debt#"
)
