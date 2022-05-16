package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk             *string `validate:"required" json:"pk"`
	Sk             *string `validate:"required" json:"sk"`
	InvestmentName *string `validate:"required" json:"investment_name"`
	InvestedAmount *int64  `validate:"required" json:"invested_amount"`
	CurrentValue   *int64  `validate:"required" json:"current_value"`
}
