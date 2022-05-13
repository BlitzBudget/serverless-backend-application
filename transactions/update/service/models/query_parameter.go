package models

// Create struct to hold info about new item
type QueryParameter struct {
	Amount      *int64    `validate:"required" json:":a"`
	Description *string   `validate:"required" json:":d"`
	Category    *string   `validate:"required" json:":c"`
	Tags        *[]string `validate:"required" json:":t"`
	UpdatedDate *string   `validate:"required" json:":u"`
}
