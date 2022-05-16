package models

// Create struct to hold info about new item
type QueryParameter struct {
	UserId         *string `json:"userId"`
	StartsWithDate *string `json:"startsWithDate"`
	EndsWithDate   *string `json:"endsWithDate"`
	InvestedAmount *int64  `json:"invested_amount"`
	CurrentValue   *int64  `json:"current_value"`
	InvestmentName *string `json:"investment_name"`
}
