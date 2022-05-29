package models

// Create struct to hold info about retrieved item
type ResponseItem struct {
	Pk           *string `json:"pk"`
	Sk           *string `json:"sk"`
	CreationDate *string `json:"creation_date"`
	CategoryName *string `json:"category_name"`
	CategoryType *string `json:"category_type"`
}

type ResponseItems []*ResponseItem
