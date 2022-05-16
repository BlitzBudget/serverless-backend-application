package models

// Create struct to hold info about new item
type QueryParameter struct {
	InvestmentName *string `validate:"required" json:":n"`
	InvestedAmount *int64  `validate:"required" json:":a"`
	CurrentValue   *int64  `validate:"required" json:":c"`
	UpdatedDate    *string `validate:"required" json:":u"`
}
