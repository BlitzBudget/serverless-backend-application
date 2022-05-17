package models

// Create struct to hold info about new item
type QueryParameter struct {
	CategoryId *string `validate:"required" json:"categoryId"`
}
