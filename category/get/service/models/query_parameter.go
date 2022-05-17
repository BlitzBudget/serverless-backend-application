package models

// Create struct to hold info about new item
type QueryParameter struct {
	UserId         *string `validate:"required" json:"userId"`
	StartsWithDate *string `validate:"required" json:"startsWithDate"`
	EndsWithDate   *string `validate:"required" json:"endsWithDate"`
}
