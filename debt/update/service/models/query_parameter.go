package models

// Create struct to hold info about new item
type QueryParameter struct {
	DebtedAmount *int64  `validate:"required" json:":a"`
	DebtName     *string `validate:"required" json:":n"`
	CurrentValue *int64  `validate:"required" json:":c"`
	UpdatedDate  *string `validate:"required" json:":u"`
	DebtRepaid   *bool   `validate:"required" json:":r"`
}
