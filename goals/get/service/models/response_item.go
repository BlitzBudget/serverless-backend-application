package models

// Create struct to hold info about retrieved item
type ResponseItem struct {
	Pk           *string `json:"pk"`
	Sk           *string `json:"sk"`
	CreationDate *string `json:"creation_date"`
	TargetAmount *int64  `json:"target_amount"`
	TargetDate   *string `json:"target_date"`
	Name         *string `json:"name"`
}

type ResponseItems []*ResponseItem
