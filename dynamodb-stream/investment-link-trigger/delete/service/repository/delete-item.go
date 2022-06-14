package repository

import (
	"add-debt-link/service/config"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Item has the values for a movie
type item struct {
	Pk *string `json:"pk"`
	Sk *string `json:"sk"`
}

func DeleteItem(pk *string, sk *string, svc dynamodbiface.DynamoDBAPI) {
	item := item{
		Pk: pk,
		Sk: sk,
	}

	fmt.Printf("Deleting Item %v with the PK %v. \n", sk, pk)
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return
	}

	input := &dynamodb.DeleteItemInput{
		Key:       av,
		TableName: &config.TableName,
	}

	_, err = svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Successfully deleted the Debt rule %v for Debt %v. \n", sk, pk)
}
