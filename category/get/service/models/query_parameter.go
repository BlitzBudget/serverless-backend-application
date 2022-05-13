package models

// Create struct to hold info about new item
type QueryParameter struct {
	UserId         *string `json:"userId"`
	StartsWithDate *string `json:"startsWithDate"`
	EndsWithDate   *string `json:"endsWithDate"`
}
