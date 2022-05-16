package models

// Create struct to hold info about new item
type QueryParameter struct {
	Currency    *string `validate:"required" json:":c"`
	Name        *string `validate:"required" json:":n"`
	UpdatedDate *string `validate:"required" json:":u"`
}
