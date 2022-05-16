package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk                 *string `json:"pk"`
	Sk                 string  `json:"sk"`
	CurrentValue       *int64  `json:"current_value"`
	InvestmentedAmount *int64  `json:"investmented_amount"`
	CreationDate       *string `json:"creation_date"`
	UpdatedDate        *string `json:"updated_date"`
	InvestmentName     *string `json:"investment"`
}
