package models

// Create struct to hold info about retrieved item
type ResponseItem struct {
	Pk           *string `json:"pk"`
	Sk           *string `json:"sk"`
	CreationDate *string `json:"creation_date"`
	TagName      *string `json:"tag_name"`
}

type ResponseItems []*ResponseItem
