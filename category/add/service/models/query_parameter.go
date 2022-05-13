package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string `json:"pk"`
	Sk           string  `json:"sk"`
	CategoryType *string `json:"category_type"`
	CreationDate *string `json:"creation_date"`
	UpdatedDate  *string `json:"updated_date"`
	CategoryName *string `json:"category"`
}
