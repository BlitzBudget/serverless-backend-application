package repository

import (
	"fmt"
	"get-transactions/service/config"
	"get-transactions/service/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func QueryItem(av *models.QueryParameter, svc *dynamodb.DynamoDB) (*dynamodb.QueryOutput, error) {
	startsWithDate := config.SkPrefix + *av.StartsWithDate
	endsWithDate := config.SkPrefix + *av.EndsWithDate
	fmt.Printf("The starts With Date is %v and the Ends with Date is %v", startsWithDate, endsWithDate)

	input := &dynamodb.QueryInput{
		TableName: aws.String(config.TableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(*av.WalletId),
					},
				},
			},
			"sk": {
				ComparisonOperator: aws.String("BETWEEN"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(startsWithDate),
					},
					{
						S: aws.String(endsWithDate),
					},
				},
			},
		},
		ProjectionExpression: &config.ProjectionExpression,
		ScanIndexForward:     &config.ScanIndexForward,
	}

	data, err := svc.Query(input)
	return data, err
}
