package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string `json:"pk"`
	Sk           string  `json:"sk"`
	TargetAmount *int64  `json:"target_amount"`
	TargetDate   *string `json:"target_date"`
	Name         *string `json:"name"`
	CreationDate *string `json:"creation_date"`
	UpdatedDate  *string `json:"updated_date"`
}
