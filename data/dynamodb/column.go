package dynamostore

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/driabwb/retroboard/domain"
)

type ColumnDynamoStore struct {
	client DynamoDB
}

type dynamoColumn struct {
	dynamoKey
	domain.Column
}

func makeColumnHK(boardID string) string {
	return key(boardKey, boardID)
}

func makeColumnSK(id string) string {
	return key(columnkey, id)
}

func makeDynamoColumn(col domain.Column) dynamoColumn {
	return dynamoColumn{
		dynamoKey: dynamoKey{
			HK: makeColumnHK(col.BoardID),
			SK: makeColumnSK(col.ID),
		},
		Column: col,
	}
}

func makeColumnKey(id, boardID string) (map[string]types.AttributeValue, error) {
	key := dynamoKey{
		HK: makeColumnHK(BoardID),
		SK: makeColumnSK(id),
	}
	return attributevalue.Marshal(key)
}

func NewColumnDynamoStore(client DynamoDB) ColumnDynamoStore {
	return ColumnDynamoStore{
		client: client,
	}
}

func (cds ColumnDynamoStore) SaveColumn(ctx context.Context, col domain.Column) error {
	dCol := makeDynamoColumn(col)
	mCol, err := attributevalue.Marshal(dCol)
	if err != nil {
		return fmt.Errorf("Failed to Marshal column: %w", err)
	}

	in := &dynamodb.PutItemInput{
		TableName: TableName,
		Item:      mCol,
	}

	_, err = cds.client.PutItem(ctx, in)
	if err != nil {
		return fmt.Errorf("Failed to write column: %w", err)
	}

	return nil
}

func (cds ColumnDynamoStore) DeleteColumn(ctx context.Context, id, boardID string) error {
	key, err := makeColumnKey(id, boardID)
	if err != nil {
		return fmt.Errorf("Failed to Marshal column key: %w", err)
	}

	in := &dynamodb.DeleteItemInput{
		TableName: TableName,
		Key:       key,
	}

	_, err = cds.client.DeleteItem(ctx, in)
	if err != nil {
		return fmt.Errorf("Failed to delete column: %w", err)
	}

	return nil
}

func (cds ColumnDynamoStore) UpdateTitle(ctx context.Context, id, boardID, title string) error {
	key, err := makeColumnKey(id, boardID)
	if err != nil {
		return fmt.Errorf("Failed to Marshal column key: %w", err)
	}

	update := expression.Set(expression.Name(titleProp), expression.Value(title))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return fmt.Errorf("Failed to build update expression: %w", err)
	}

	in := &dynamodb.UpdateItemInput{
		TableName:                 TableName,
		Key:                       key,
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	_, err = cds.client.UpdateItem(ctx, in)
	if err != nil {
		return fmt.Errorf("Failed to update item: %w", err)
	}

	return nil
}
