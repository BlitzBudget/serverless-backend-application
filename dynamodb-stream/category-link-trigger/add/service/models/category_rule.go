package models

// Create struct to hold info about retrieved item
type CategoryRule struct {
	Pk              *string `json:"pk"`
	Sk              *string `json:"sk"`
	CreationDate    *string `json:"creation_date"`
	TransactionName *string `json:"transaction_name"`
	CategoryId      *string `json:"category_id"`
}