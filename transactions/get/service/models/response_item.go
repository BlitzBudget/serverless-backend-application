package models

// Create struct to hold info about retrieved item
type ResponseItem struct {
	Pk           *string   `json:"pk"`
	Sk           *string   `json:"sk"`
	CreationDate *string   `json:"creation_date"`
	CategoryId   *string   `json:"category_id"`
	Description  *string   `json:"description"`
	Amount       *int64    `json:"amount"`
	Tags         *[]string `json:"tags"`
	InvestmentId *string   `json:"investment_id"`
	DebtId       *string   `json:"debt_id"`
	GoalId       *string   `json:"goal_id"`
}

type ResponseItems []*ResponseItem
