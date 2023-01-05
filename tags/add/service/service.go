package service

import (
	getTagService "add-tag/get-tag/service"
	"add-tag/service/repository"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func SaveRequest(body *string) error {
	// snippet-start:[dynamodb.go.create_item.session]
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	// snippet-end:[dynamodb.go.create_item.session]

	queryParameter, err := repository.TransformToQueryParamter(body)
	if err != nil {
		fmt.Printf("SaveRequest: Got error marshalling new item: %v \n", err)
		return err
	}
	responseItems := getTagService.QueryItems(queryParameter)

	for _, item := range responseItems {
		if item.TagName == queryParameter.TagName {
			err = fmt.Errorf("SaveRequest: TagName already exists")
			return err
		}
	}

	err = repository.MandatoryFieldsCheck(queryParameter)
	if err != nil {
		fmt.Printf("SaveRequest: Mandatory Field Check Failure: %v \n", err)
		return err
	}

	av, err := repository.AttributeBuilder(queryParameter)
	if err != nil {
		fmt.Printf("SaveRequest: Attribute Builder error: %v \n", err)
		return err
	}

	err = repository.CreateItem(av, svc)
	if err != nil {
		fmt.Printf("SaveRequest: Error Creating Item: %v \n", err)
		return err
	}

	fmt.Println("Successfully added the item!")
	return nil
}
