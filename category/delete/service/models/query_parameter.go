package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string   `json:"pk"`
	Sk           string    `json:"sk"`
	Amount       *int64    `json:"amount"`
	Description  *string   `json:"description"`
	CreationDate *string   `json:"creation_date"`
	UpdatedDate  *string   `json:"updated_date"`
	Category     *string   `json:"category"`
	Tags         *[]string `json:"tags"`
}
