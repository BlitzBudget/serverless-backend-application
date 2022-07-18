package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string   `validate:"required" json:"pk"`
	Sk           string    `json:"sk"`
	Amount       *float64  `validate:"required" json:"amount"`
	Description  *string   `validate:"required" json:"description"`
	CreationDate *string   `json:"creation_date"`
	UpdatedDate  *string   `json:"updated_date"`
	CategoryId   *string   `validate:"required" json:"category_id"`
	Tags         *[]string `validate:"required" json:"tags"`
}
