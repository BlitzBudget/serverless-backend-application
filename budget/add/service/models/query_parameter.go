package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string  `validate:"required" json:"pk"`
	Sk           string   `json:"sk"`
	Planned      *float64 `validate:"required" json:"planned"`
	CreationDate *string  `json:"creation_date"`
	UpdatedDate  *string  `json:"updated_date"`
	Category     *string  `validate:"required" json:"category"`
}
