package helper

import (
	"delete-category-link/service/config"
	"delete-category-link/service/models"
	"delete-category-link/service/repository"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func FetchWallets(svc dynamodbiface.DynamoDBAPI, category *models.Category) []*models.Wallet {
	queryOutput, err := repository.QueryItems(svc, category.Pk, config.SkWalletPrefix, config.WalletProjectionExpression)
	if err != nil {
		fmt.Printf("FetchWallets: Got error while Fetching Wallet: %v. \n", err)
		return nil
	}

	var wallets []*models.Wallet
	wallets, err = parseWalletResponse(queryOutput)
	if err != nil {
		fmt.Printf("RemoveCategoryLink: Got error Category Rule ParseResponse: %v. \n", err)
		return nil
	}

	return wallets
}

func parseWalletResponse(result *dynamodb.QueryOutput) ([]*models.Wallet, error) {

	if result.Items == nil {
		msg := "no Items found in Wallets"
		return nil, errors.New(msg)
	}

	wallets := []*models.Wallet{}
	var err error

	for k, v := range result.Items {
		wallet := models.Wallet{}

		err = dynamodbattribute.UnmarshalMap(v, &wallet)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record %v, %v \n", k, err))
		}
		wallets = append(wallets, &wallet)
	}

	fmt.Printf("Parsed %v Items for wallets. \n", len(wallets))
	return wallets, nil
}
