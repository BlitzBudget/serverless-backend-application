package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk             *string  `validate:"required" json:"pk"`
	Sk             *string  `validate:"required" json:"sk"`
	TagName        *string  `validate:"required" json:"tag_name"`
}
