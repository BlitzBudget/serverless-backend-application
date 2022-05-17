package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk           *string `validate:"required" json:"pk"`
	Sk           *string `validate:"required" json:"sk"`
	Debt         *int64  `validate:"required" json:"debt"`
	CurrentValue *int64  `validate:"required" json:"current_value"`
	DebtName     *string `validate:"required" json:"debt_name"`
	DebtRepaid   *bool   `validate:"required" json:"debt_repaid"`
}