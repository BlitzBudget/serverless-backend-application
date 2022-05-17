package models

// Create struct to hold info about new item
type RequestModel struct {
	Pk            *string `validate:"required" json:"pk"`
	Sk            *string `validate:"required" json:"sk"`
	TargetAmount  *int64  `validate:"required" json:"target_amount"`
	GoalAchieved  *bool   `validate:"required" json:"goal_achieved"`
	CurrentAmount *int64  `validate:"required" json:"current_amount"`
	TargetDate    *string `validate:"required" json:"target_date"`
	Name          *string `validate:"required" json:"name"`
}
