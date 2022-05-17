package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "goal_achieved, target_amount, target_date, name, sk, pk, creation_date"

const (
	SkPrefix = "Goal#"
)
