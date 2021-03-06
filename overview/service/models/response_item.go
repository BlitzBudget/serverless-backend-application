package models

// Create struct to hold info about retrieved item
type ResponseItem struct {
	Pk                   *string   `json:"pk"`
	Sk                   *string   `json:"sk"`
	CreationDate         *string   `json:"creation_date"`
	Category             *string   `json:"category"`
	Description          *string   `json:"description"`
	Amount               *float64  `json:"amount"`
	Tags                 *[]string `json:"tags"`
	ScheduledTransaction *bool     `json:"scheduled_transaction"`
}

type ResponseItems []*ResponseItem
