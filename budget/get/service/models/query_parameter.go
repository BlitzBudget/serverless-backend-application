package models

// Create struct to hold info about new item
type QueryParameter struct {
	WalletId       *string `validate:"required" json:"wallet_id"`
	StartsWithDate *string `validate:"required" json:"starts_with_date"`
	EndsWithDate   *string `validate:"required" json:"ends_with_date"`
}
