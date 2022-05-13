package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk           *string `validate:"required" json:"pk"`
	Sk           *string `validate:"required" json:"sk"`
	TargetAmount *int64  `validate:"required" json:"target_amount"`
	TargetDate   *string `validate:"required" json:"target_date"`
	Name         *string `validate:"required" json:"name"`
}
