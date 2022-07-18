package models

// Create struct to hold info about retrieved item
type GoalRule struct {
	Pk                *string  `json:"pk"`
	Sk                *string  `json:"sk"`
	CreationDate      *string  `json:"creation_date"`
	TransactionName   *string  `json:"transaction_name"`
	TransactionAmount *float64 `json:"transaction_amount"`
	GoalId            *string  `json:"goal_id"`
}
