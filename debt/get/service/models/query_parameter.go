package models

// Create struct to hold info about new item
type QueryParameter struct {
	UserId *string `validate:"required" json:"userId"`
}
