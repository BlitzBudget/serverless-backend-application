package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string `validate:"required" json:"pk"`
	Sk           string  `json:"sk"`
	CreationDate *string `json:"creation_date"`
	UpdatedDate  *string `json:"updated_date"`
	TagName      *string `validate:"required" json:"tag_name"`
}
