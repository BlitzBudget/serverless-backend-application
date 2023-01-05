package models

// Create struct to hold info about new item
type QueryParameter struct {
	TagName        *string  `validate:"required" json:":n"`
	UpdatedDate    *string  `validate:"required" json:":u"`
}
