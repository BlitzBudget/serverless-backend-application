package models

// Create struct to hold info about new item
type QueryParameter struct {
	DebtId *string `validate:"required" json:"debtId"`
}
