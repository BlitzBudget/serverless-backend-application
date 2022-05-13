package models

// Create struct to hold info about new item
type QueryParameter struct {
	Planned     *int64  `validate:"required" json:":p"`
	Category    *string `validate:"required" json:":c"`
	UpdatedDate *string `validate:"required" json:":u"`
}
