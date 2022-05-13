package models

// Create struct to hold info about new item
type QueryParameter struct {
	CategoryName *string `validate:"required" json:":n"`
	CategoryType *string `validate:"required" json:":t"`
	UpdatedDate  *string `validate:"required" json:":u"`
}
