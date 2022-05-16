package models

// Create struct to hold info about retrieved item
type ResponseItem struct {
	Pk           *string `json:"pk"`
	Sk           *string `json:"sk"`
	CreationDate *string `json:"creation_date"`
	Currency     *string `json:"currency"`
	Name         *string `json:"name"`
}

type ResponseItems []*ResponseItem
