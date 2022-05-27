package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk              *string `validate:"required" json:"pk"`
	Sk              string  `json:"sk"`
	TransactionName *string `validate:"required" json:"transaction_name"`
	CategoryId      *string `validate:"required" json:"category_id"`
	CreationDate    *string `json:"creation_date"`
	UpdatedDate     *string `json:"updated_date"`
}
