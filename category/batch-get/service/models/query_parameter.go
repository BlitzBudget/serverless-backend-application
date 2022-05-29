package models

// Create struct to hold info about new item
type QueryParameter struct {
	UserId      *string   `validate:"required" json:"user_id"`
	CategoryIDs *[]string `validate:"required" json:"category_ids"`
}
