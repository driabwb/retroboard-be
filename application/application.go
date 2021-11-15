package application

import (
	"context"

	"github.com/driabwb/retroboard/domain"
)

type (
	CardCreator interface {
		CreateCard(ctx context.Context, card domain.Card) (domain.Card, error)
	}

	CardContentUpdater interface {
		UpdateCardContent(ctx context.Context, cardID, boardID, content string) error
	}

	CardColumnUpdater interface {
		UpdateCardColumn(ctx context.Context, cardID, columnID, boardID string) error
	}

	CardVoteUpdater interface {
		UpdateCardVotes(ctx context.Context, cardID, boardID string, delta int) error
	}

	CardDeleter interface {
		DeleteCard(ctx context.Context, cardID, boardID string) error
	}
)

type (
	ColumnCreator interface {
		CreateColumn(ctx context.Context, col domain.Column) (domain.Column, error)
	}

	ColumnTitleUpdater interface {
		UpdateColumnTitle(ctx context.Context, columnID, boardID, title string) error
	}

	ColumnDeleter interface {
		DeleteColumn(ctx context.Context, columnID, boardID string) error
	}
)

type (
	BoardCreator interface {
		CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error)
	}

	BoardTitleUpdater interface {
		UpdateBoardTitle(ctx context.Context, boardID, title string) error
	}
)

type (
	BoardGetter interface {
		GetBoard(ctx context.Context, boardID string) (domain.Board, error)
	}
)

type (
	App struct {
		datastore RetroboardDatastore
	}
)

func NewApp(datastore RetroboardDatastore) App {
	return App{
		datastore: datastore,
	}
}
