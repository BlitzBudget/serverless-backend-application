package models

// Create struct to hold info about new item
type QueryParameter struct {
	Sk              string  `validate:"required" json:":s"`
	TransactionName *string `validate:"required" json:":t"`
	UpdatedDate     *string `validate:"required" json:":u"`
}
