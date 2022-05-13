package models

// Create struct to hold info about new item
type QueryParameter struct {
	TargetAmount *int64  `validate:"required" json:":a"`
	TargetDate   *string `validate:"required" json:":d"`
	Name         *string `validate:"required" json:":n"`
	UpdatedDate  *string `validate:"required" json:":u"`
}
