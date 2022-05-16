package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk       *string `validate:"required" json:"pk"`
	Sk       *string `validate:"required" json:"sk"`
	Currency *string `validate:"required" json:"currency"`
	Name     *string `validate:"required" json:"name"`
}
