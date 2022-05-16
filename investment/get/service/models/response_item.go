package models

// Create struct to hold info about retrieved item
type ResponseItem struct {
	Pk             *string `json:"pk"`
	Sk             *string `json:"sk"`
	CreationDate   *string `json:"creation_date"`
	InvestedAmount *int64  `json:"invested_amount"`
	CurrentValue   *int64  `json:"current_value"`
	InvestmentName *string `json:"investment_name"`
}

type ResponseItems []*ResponseItem
