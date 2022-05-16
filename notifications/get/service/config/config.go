package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "notification, sk, pk, creation_date"

const (
	SkPrefix = "Notification#"
)
