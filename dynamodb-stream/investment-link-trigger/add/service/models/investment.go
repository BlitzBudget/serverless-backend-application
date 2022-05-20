package models

// Create struct to hold info about retrieved item
type Investment struct {
	Pk                 *string `json:"pk"`
	Sk                 *string `json:"sk"`
	CreationDate       *string `json:"creation_date"`
	InvestmentedAmount *int64  `json:"investmented_amount"`
	CurrentValue       *int64  `json:"current_value"`
	InvestmentName     *string `json:"investment_name"`
	InvestmentAmount   *bool   `json:"investment_amount"`
}
