package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string `validate:"required" json:"pk"`
	Sk           string  `validate:"required" json:"sk"`
	Notification *string `validate:"required" json:"notification"`
	CreationDate *string `json:"creation_date"`
	UpdatedDate  *string `json:"updated_date"`
}
