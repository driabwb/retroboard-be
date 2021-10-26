package dynamostore

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/driabwb/retroboard/domain"
)

type EntityDynamoStore struct {
	client DynamoDB
}

func NewEnitiyDynamoStore(client DynamoDB) EntityDynamoStore {
	return EntityDynamoStore{
		client: client,
	}
}

func (eds EntityDynamoStore) GetEntity(ctx context.Context, boardID string) (Entity, error) {
	hk := expression.Key(hashKeyProp).Equal(expression.Value(makeBoardHK(boardID)))
	expr, err := expression.NewBuilder().WithKeyCondition(hk).Build()
	if err != nil {
		return nil, fmt.Errorf("Failed to build entity query: %w", err)
	}

	in := &dynamodb.QueryInput{
		TableName:                 TableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	var board domain.Board
	cols := make([]domain.Column, 0)
	cards := make([]domain.Card, 0)

	for {
		result, err := eds.client.Query(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("Error querying for entity data: %w", err)
		}

		for _, elem := range result.Items {
			switch itemDynamoType(elem) {
			case dynamoBoardItem:
				if board.ID != "" {
					return nil, fmt.Errorf("Found two board entries in dynamo: %s", boardID)
				}
				err = attributevalue.UnmarshalMap(elem, &board)
				if err != nil {
					return nil, fmt.Errorf("Failed to unmarshal board type: %w", err)
				}
			case dynamoColumnItem:
				col := domain.Column{}
				err = attributevalue.UnmarshalMap(elem, &col)
				if err != nil {
					return nil, fmt.Errorf("Failed to unmarshal column %w", err)
				}
				cols = append(cols, col)
			case dynamoCardItem:
				card := domain.Card{}
				err = attributevalue.UnmarshalMap(elem, &card)
				if err != nil {
					return nil, fmt.Errorf("Failed to unmarshal card %w", err)
				}
				cards = append(cards, card)
			default:
				// Do nothing for unknown item types
			}
		}

		if result.LastEvaluatedKey == nil {
			break
		}

		in.ExclusiveStartKey = result.LastEvaluatedKey
	}

	return Entity{
		Board: board,
		Cols:  cols,
		Cards: cards,
	}
}

func itemDynamoType(item map[string]types.AttributeValue) string {
	sortKey, ok := item[sortKeyProp]
	if !ok {
		return "Unknown"
	}

	itemTypePrefix := strings.Split(sortKey, keySeparator)[0]

	switch itemTypePrefix {
	case boardKey:
		return dynamoBoardItem
	case columnKey:
		return dynamoColumnItem
	case cardKey:
		return dynamoCardItem
	default:
		return dynamoUnknownItem
	}
}
