package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string `json:"pk"`
	Sk           string  `json:"sk"`
	Notification *string `json:"notification"`
	CreationDate *string `json:"creation_date"`
	UpdatedDate  *string `json:"updated_date"`
}