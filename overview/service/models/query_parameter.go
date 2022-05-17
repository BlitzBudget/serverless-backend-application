package models

// Create struct to hold info about new item
type QueryParameter struct {
	WalletId       *string `validate:"required" json:"walletId"`
	StartsWithDate *string `validate:"required" json:"startsWithDate"`
	EndsWithDate   *string `validate:"required" json:"endsWithDate"`
}
