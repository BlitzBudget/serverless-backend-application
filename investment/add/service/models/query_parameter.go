package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk             *string  `validate:"required" json:"pk"`
	Sk             string   `json:"sk"`
	CurrentValue   *float64 `json:"current_value"`
	InvestedAmount *float64 `validate:"required" json:"invested_amount"`
	CreationDate   *string  `json:"creation_date"`
	UpdatedDate    *string  `json:"updated_date"`
	InvestmentName *string  `validate:"required" json:"investment_name"`
}
