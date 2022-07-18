package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk       *string  `validate:"required" json:"pk"`
	Sk       *string  `validate:"required" json:"sk"`
	Planned  *float64 `validate:"required" json:"planned"`
	Category *string  `validate:"required" json:"category"`
}
