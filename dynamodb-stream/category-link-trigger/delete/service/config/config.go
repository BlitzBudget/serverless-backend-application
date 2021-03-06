package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "category_id, transaction_amount, transaction_name, sk, pk, creation_date"

const (
	SkCategoryRulePrefix       = "CategoryRule#"
	SkWalletPrefix             = "Wallet#"
	WalletProjectionExpression = "wallet_currency, wallet_name, sk, pk, creation_date"
)
