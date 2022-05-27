package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk            *string `validate:"required" json:"pk"`
	Sk            string  `json:"sk"`
	TargetAmount  *int64  `validate:"required" json:"target_amount"`
	CurrentAmount *int64  `json:"current_amount"`
	TargetDate    *string `validate:"required" json:"target_date"`
	Name          *string `validate:"required" json:"goal_name"`
	GoalAchieved  *bool   `json:"goal_achieved"`
	CreationDate  *string `json:"creation_date"`
	UpdatedDate   *string `json:"updated_date"`
}
