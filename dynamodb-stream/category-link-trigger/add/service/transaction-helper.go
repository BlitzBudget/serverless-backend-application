package service

import (
	"add-category-link/service/models"
	"add-category-link/service/repository"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type QueryParameter struct {
	CategoryId  *string `validate:"required" json:":d"`
	UpdatedDate *string `validate:"required" json:":u"`
}

// convert item to dynamodb attribute
func ParseToQueryParameter(categoryId *string) map[string]*dynamodb.AttributeValue {
	date := time.Now().Format(time.RFC3339Nano)
	av, err := dynamodbattribute.MarshalMap(QueryParameter{
		CategoryId:  categoryId,
		UpdatedDate: &date,
	})

	if err != nil {
		fmt.Printf("ParseToQueryParameter: Failed to update transaction with category id %v, %v. \n", *categoryId, err)
	}

	fmt.Printf("Successfully converted the update transaction request for the category id %v \n", *categoryId)
	return av
}

// Update Transaction Item with Category ID
func UpdateTransactionWithCategoryId(categoryId *string, svc dynamodbiface.DynamoDBAPI, request *models.Transaction) {
	fmt.Printf("Updating the transaction %v with the category id %v \n", *request.Sk, *categoryId)

	av := ParseToQueryParameter(categoryId)
	updateExpression := "set category_id = :d, updated_date = :u"
	repository.UpdateItem(av, svc, request.Pk, request.Sk, updateExpression)
}
