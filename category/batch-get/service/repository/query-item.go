package repository

import (
	"batch-get-category/service/config"
	"batch-get-category/service/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func QueryItem(av *models.QueryParameter, svc dynamodbiface.DynamoDBAPI) (*dynamodb.QueryOutput, error) {
	input := &dynamodb.QueryInput{
		TableName: aws.String(config.TableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(*av.UserId),
					},
				},
			},
			"sk": {
				ComparisonOperator: aws.String("BEGINS_WITH"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(config.SkPrefix),
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
