package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk           *string `validate:"required" json:"pk"`
	Sk           *string `validate:"required" json:"sk"`
	CategoryName *string `validate:"required" json:"category_name"`
	CategoryType *string `validate:"required" json:"category_type"`
}
