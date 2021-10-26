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

type BoardDynamoStore struct {
	client DynamoDB
}

type dynamoBoard struct {
	dynamoKey
	domain.Board
}

func makeBoardHK(boardID string) string {
	return key(boardKey, boardID)
}

func makeBoardSK() string {
	return boardKey
}

func makeDynamoBoard(board domain.Board) dynamoBoard {
	return dynamoBoard{
		dynamoKey: dynamoKey{
			HK: makeBoardHK(board.ID),
			SK: makeBoardSK(),
		},
		Board: board,
	}
}

func makeBoardKey(id string) (map[string]types.AttributeValue, error) {
	key := dynamoKey{
		HK: makeBoardK(col.BoardID),
		SK: makeBoardK(),
	}
	return attributevalue.Marshal(key)
}

func NewBoardDynamoStore(client DynamoDB) BoardDynamoStore {
	return BoardDynamoStore{
		client: client,
	}
}

func (bds BoardDynamoStore) SaveBoard(ctx context.Context, board domain.Board) error {
	dBoard := makeDynamoBoard(board)
	mBoard, err := attributevalue.Marshal(dBoard)
	if err != nil {
		return fmt.Errorf("Failed to marshal board: %w", err)
	}

	in := &dynamodb.PutItemInput{
		TableName: TableName,
		Item:      mBoard,
	}

	_, err = cds.client.PutItem(ctx, in)
	if err != nil {
		return fmt.Errorf("Failed to write board: %w", err)
	}
	return nil
}

func (bds BoardDynamoStore) UpdateTitle(ctx context.Context, id, title string) error {
	key, err := makeBoardKey(id)
	if err != nil {
		return fmt.Errorf("Failed to Marshal board key: %w", err)
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
