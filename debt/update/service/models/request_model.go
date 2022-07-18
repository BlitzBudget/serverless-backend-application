package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk           *string  `validate:"required" json:"pk"`
	Sk           *string  `validate:"required" json:"sk"`
	DebtedAmount *float64 `validate:"required" json:"debted_amount"`
	CurrentValue *float64 `validate:"required" json:"current_value"`
	DebtName     *string  `validate:"required" json:"debt_name"`
	DebtRepaid   *bool    `validate:"required" json:"debt_repaid"`
}
