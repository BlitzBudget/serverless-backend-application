package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk          *string   `validate:"required" json:"pk"`
	Sk          *string   `validate:"required" json:"sk"`
	Amount      *int64    `validate:"required" json:"amount"`
	Description *string   `validate:"required" json:"description"`
	Category    *string   `validate:"required" json:"category"`
	Tags        *[]string `validate:"required" json:"tags"`
}
