package models

// Create struct to hold info about new item
type QueryParameter struct {
	WalletId       *string `json:"walletId"`
	StartsWithDate *string `json:"startsWithDate"`
	EndsWithDate   *string `json:"endsWithDate"`
}
