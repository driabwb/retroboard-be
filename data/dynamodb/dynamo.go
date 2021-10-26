package dynamostore

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDB interface {
	DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
}

type dynamoKey struct {
	HK string
	SK string
}

type dynamoGSI1Key struct {
	GSI1HK string
	GSI1SK string
}

const (
	// Single Table
	TableName = "Retroboard"

	// Key Prefixes
	boardKey  = "board"
	columnKey = "col"
	cardKey   = "card"

	// Property Names
	hashKeyProp = "HK"
	sortKeyProp = "SK"
	gsi1HKProp  = "GSI1HK"
	gsi1SKProp  = "GSI1SK"
	columnProp  = "ColumnID"
	contentProp = "Content"
	votesProp   = "Votes"
	titleProp   = "Title"

	// Dynamo Item Types
	dynamoBoardItem   = "dynamoBoardItem"
	dynamoColumnItem  = "dynamoColumnItem"
	dynamoCardItem    = "dynamoCardItem"
	dynamoUnknownItem = "dynamoUnknownItem"

	// Dynamo Key Separator
	keySeparator = ":"
)

func key(parts ...string) string {
	return strings.Join(parts, keySeparator)
}
