package datastore

import (
	"context"

	"github.com/driabwb/retroboard/domain"
)

type CardStorer interface {
	SaveCard(ctx context.Context, card domain.Card) error
}

type CardColumnUpdater interface {
	UpdateColumn(ctx context.Context, id, columnID, boardID string) error
}

type CardContentUpdater interface {
	UpdateContent(ctx context.Context, id, boardID, content string) error
}

type CardVoteUpdater interface {
	AddVotes(ctx context.Context, id, boardID string, delta int) error
}

type CardDeleter interface {
	DeleteCard(ctx context.Context, id, boardID string) error
}
