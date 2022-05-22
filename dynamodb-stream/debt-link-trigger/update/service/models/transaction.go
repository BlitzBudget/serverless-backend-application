package models

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create struct to hold info about retrieved item
type Transaction struct {
	Pk                   *string   `json:"pk"`
	Sk                   *string   `json:"sk"`
	CreationDate         *string   `json:"creation_date"`
	Category             *string   `json:"category"`
	Description          *string   `json:"description"`
	Amount               *int64    `json:"amount"`
	Tags                 *[]string `json:"tags"`
	ScheduledTransaction *bool     `json:"scheduled_transaction"`
}

func ConvertDynamoDBToModel(dbAttrMap map[string]*dynamodb.AttributeValue) Transaction {
	amount := dbAttrMap["amount"].N
	amt, err := strconv.Atoi(*amount)
	if err != nil {
		fmt.Printf("AttributeBuilder: Unable to convert amount to integer: %v. \n", err)
		amt = 0
	}
	bigIntAmount := int64(amt)

	transaction := Transaction{
		Pk:          dbAttrMap["pk"].S,
		Sk:          dbAttrMap["sk"].S,
		Description: dbAttrMap["description"].S,
		Amount:      &bigIntAmount,
		Category:    dbAttrMap["category"].S,
	}
	return transaction
}
