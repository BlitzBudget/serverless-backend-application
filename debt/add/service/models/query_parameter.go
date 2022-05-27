package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string `validate:"required" json:"pk"`
	Sk           string  `json:"sk"`
	CurrentValue *int64  `validate:"required" json:"current_value"`
	DebtedAmount *int64  `validate:"required" json:"debted_amount"`
	CreationDate *string `json:"creation_date"`
	UpdatedDate  *string `json:"updated_date"`
	DebtRepaid   *bool   `json:"debt_repaid"`
	DebtName     *string `validate:"required" json:"debt_name"`
}
