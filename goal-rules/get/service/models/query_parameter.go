package models

// Create struct to hold info about new item
type QueryParameter struct {
	GoalId *string `validate:"required" json:"goalId"`
}
