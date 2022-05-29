package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "category_name, category_type, sk, pk, creation_date"

const (
	SkPrefix = "Category#"
)
