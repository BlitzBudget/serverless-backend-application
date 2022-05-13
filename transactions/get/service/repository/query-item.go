package repository

import (
	"get-transactions/service/config"
	"get-transactions/service/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func QueryItem(av *models.QueryParameter, svc *dynamodb.DynamoDB) (*dynamodb.QueryOutput, error) {
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
						S: aws.String(config.SkPrefix + *av.StartsWithDate),
					},
					{
						S: aws.String(config.SkPrefix + *av.EndsWithDate),
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
