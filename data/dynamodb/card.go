package dynamostore

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/driabwb/retroboard/domain"
)

type CardDynamoStore struct {
	client DynamoDB
}

type dynamoCard struct {
	dynamoKey
	dynamoGSI1Key
	domain.Card
}

func makeCardHK(boardID string) string {
	return key(boardKey, boardID)
}

func makeCardSK(id string) string {
	return key(cardKey, id)
}

func makeCardGSI1HK(colID string) string {
	return key(columnKey, colID)
}

func makeCardGSI1SK(id string) string {
	return key(cardKey, id)
}

func makeDynamoCard(card domain.Card) dynamoCard {
	return dynamoCard{
		dynamoKey: dynamoKey{
			HK: makeCardHK(card.BoardID),
			SK: makeCardSK(card.ID),
		},
		dynamoGSI1Key: dynamoCardGSI1{
			GSI1HK: makeCardGSI1HK(card.ColumnID),
			GSI1SK: makeCardGSI1SK(card.ID),
		},
		Card: card,
	}
}

func makeCardKey(id, boardID string) (map[string]types.AttributeValue, error) {
	key := dynamoKey{
		HK: makeCardHK(card),
		SK: makeCardSK(card),
	}

	return attributevalue.Marshal(key)
}

func NewCardDynamoStore(client DynamoDB) CardDynamoStore {
	return CardDynamoStore{
		client: client,
	}
}

func (cds CardDynamoStore) SaveCard(ctx context.Context, card domain.Card) error {
	dCard := makeDynamoCard(card)

	mCard, err := attributevalue.Marshal(dCard)
	if err != nil {
		return fmt.Errorf("Failed to Marshal card: %w", err)
	}

	in := &dyanmodb.PutItemInput{
		TableName: TableName,
		Item:      mCard,
	}

	_, err = cds.client.PutItem(ctx, in)
	if err != nil {
		return fmt.Errorf("Failed to write card: %w", err)
	}

	return nil
}

func (cds CardDynamoStore) DeleteCard(ctx context.Context, id, boardID string) error {
	key, err := makeCardKey(id, boardID)
	if err != nil {
		return fmt.Errorf("Failed to Marshal card key: %w", err)
	}

	in := &dyanmodb.DeleteItemInput{
		TableName: TableName,
		Key:       key,
	}

	_, err = cds.client.DeleteItem(ctx, in)
	if err != nil {
		return fmt.Errorf("Failed to delete card: %w", err)
	}

	return nil
}

func (cds CardDynamoStore) UpdateCardColumn(ctx context.Context, id, columnID, boardID string) error {
	key, err := makeCardKey(id, boardID)
	if err != nil {
		return fmt.Errorf("Failed to Marshal card key: %w", err)
	}
	update := expression.
		Set(expression.Name(columnProp), expression.Value(columnID)).
		Set(expression.Name(gsi1HKProp), expression.Value(makeCardGSI1HK(columnID)))
	cond := expression.Equal(expression.Name(hashKeyProp), expression.Value(makeCardHK(boardID)))

	expr, err := expression.NewBuilder().WithUpdate(update).WithCondition(cond).Build()
	if err != nil {
		return fmt.Errorf("Failed to build update expression: %w", err)
	}

	in := &dynamodb.UpdateItemInput{
		TableName:                 TableName,
		Key:                       key,
		UpdateExpression:          expr.Update(),
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	_, err = cds.client.UpdateItem(ctx, in)
	if err != nil {
		return fmt.Errorf("Failed to update item: %w", err)
	}

	return nil
}

func (cds CardDynamoStore) UpdateCardContent(ctx context.Context, id, boardID, content string) error {
	key, err := makeCardKey(id, boardID)
	if err != nil {
		return fmt.Errorf("Failed to Marshal card key: %w", err)
	}
	update := expression.
		Set(expression.Name(contentProp), expression.Value(content))
	cond := expression.Equal(expression.Name(hashKeyProp), expression.Value(makeCardHK(boardID)))

	expr, err := expression.NewBuilder().WithUpdate(update).WithCondition(cond).Build()
	if err != nil {
		return fmt.Errorf("Failed to build update expression: %w", err)
	}

	in := &dynamodb.UpdateItemInput{
		TableName:                 TableName,
		Key:                       key,
		UpdateExpression:          expr.Update(),
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	_, err = cds.client.UpdateItem(ctx, in)
	if err != nil {
		return fmt.Errorf("Failed to update item: %w", err)
	}
	return nil
}

func (cds CardDynamoStore) AddCardVotes(ctx context.Context, id, boardID string, delta int) error {
	key, err := makeCardKey(id, boardID)
	if err != nil {
		return fmt.Errorf("Failed to Marshal card key: %w", err)
	}
	update := expression.
		Add(expression.Name(valueProp), expression.Plus(
			expression.Name(valueProp),
			expression.Value(delta),
		),
		)
	cond := expression.Equal(expression.Name(hashKeyProp), expression.Value(makeCardHK(boardID)))

	expr, err := expression.NewBuilder().WithUpdate(update).WithCondition(cond).Build()
	if err != nil {
		return fmt.Errorf("Failed to build update expression: %w", err)
	}

	in := &dynamodb.UpdateItemInput{
		TableName:                 TableName,
		Key:                       key,
		UpdateExpression:          expr.Update(),
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	_, err = cds.client.UpdateItem(ctx, in)
	if err != nil {
		return fmt.Errorf("Failed to update item: %w", err)
	}
	return nil
}
