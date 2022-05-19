package models

// Create struct to hold info about new item
type QueryParameter struct {
	WalletId *string `validate:"required" json:"walletId"`
}
