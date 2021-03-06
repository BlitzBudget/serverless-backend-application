package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "amount, description, category_id, investment_id, debt_id, goal_id, sk, pk, creation_date, tags"

const (
	SkPrefix = "Transaction#"
)
