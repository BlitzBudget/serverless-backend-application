package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "currency, name, sk, pk, creation_date"

const (
	SkPrefix = "Wallet#"
)
