package models

// Create struct to hold info about retrieved item
type ResponseItem struct {
	Pk                *string  `json:"pk"`
	Sk                *string  `json:"sk"`
	CreationDate      *string  `json:"creation_date"`
	TransactionName   *string  `json:"transaction_name"`
	TransactionAmount *float64 `validate:"required" json:"transaction_amount"`
	InvestmentId      *string  `validate:"required" json:"investment_id"`
}

type ResponseItems []*ResponseItem
