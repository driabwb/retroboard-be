package application

import (
	"context"

	"github.com/driabwb/retroboard/domain"
	"github.com/google/uuid"
)

func (a App) CreateCard(ctx context.Context, card domain.Card) (domain.Card, error) {
	newCard := card
	newCard.ID = uuid.New()

	err := a.datastore.CreateCard(ctx, newCard)
	if err != nil {
		return domain.Card{}, err
	}

	return newCard, nil
}

func (a App) CreateColumn(ctx context.Context, col domain.Column) (domain.Column, error) {
	newCol := col
	newCol.ID = uuid.New()

	err := a.datastore.CreateColumn(ctx, newCol)
	if err != nil {
		return domain.Column{}, err
	}

	return newCol, nil
}

func (a App) CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error) {
	newBoard := board
	board.ID = uuid.New()

	err := a.datastore.CreateBoard(ctx, newCol)
	if err != nil {
		return domain.Board{}, err
	}

	return newBoard, nil
}
