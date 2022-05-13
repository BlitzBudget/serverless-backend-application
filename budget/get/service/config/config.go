package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "category, planned, sk, pk, creation_date"

const (
	SkPrefix = "Budget#"
)
