package config

import "os"

var TableName = os.Getenv("TABLE_NAME")

const (
	SkCategoryRulePrefix = "CategoryRule#"
)
