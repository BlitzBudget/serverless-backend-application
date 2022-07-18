package models

// Create struct to hold info about new item
type QueryParameter struct {
	CurrentAmount *float64 `validate:"required" json:":c"`
	TargetAmount  *float64 `validate:"required" json:":a"`
	GoalAchieved  *bool    `validate:"required" json:":g"`
	TargetDate    *string  `validate:"required" json:":d"`
	Name          *string  `validate:"required" json:":n"`
	UpdatedDate   *string  `validate:"required" json:":u"`
}
