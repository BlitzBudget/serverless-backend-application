package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk                *string  `validate:"required" json:"pk"`
	Sk                string   `json:"sk"`
	TransactionName   *string  `validate:"required" json:"transaction_name"`
	DebtId            *string  `validate:"required" json:"debt_id"`
	TransactionAmount *float64 `json:"transaction_amount"`
	CreationDate      *string  `json:"creation_date"`
	UpdatedDate       *string  `json:"updated_date"`
}
