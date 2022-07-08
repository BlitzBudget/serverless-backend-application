package models

// Create struct to hold info about retrieved item
type Wallet struct {
	Pk           *string `json:"pk"`
	Sk           *string `json:"sk"`
	CreationDate *string `json:"creation_date"`
	Currency     *string `json:"wallet_currency"`
	Name         *string `json:"wallet_name"`
}

type Wallets []*Wallet
