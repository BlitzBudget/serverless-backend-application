package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "goal_achieved, target_amount, current_amount, target_date, goal_name, sk, pk, creation_date"

const (
	SkPrefix = "Goal#"
)
