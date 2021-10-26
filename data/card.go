package datastore

import (
	"context"

	"github.com/driabwb/retroboard/domain"
)

type CardStorer interface {
	SaveCard(ctx context.Context, card domain.Card) error
}

type CardUpdater interface {
	UpdateColumn(ctx context.Context, id, columnID, boardID string) error
	UpdateContent(ctx context.Context, id, boardID, content string) error
	AddVotes(ctx context.Context, id, boardID string, delta int) error
}

type CardDeleter interface {
	DeleteCard(ctx context.Context, id, boardID string) error
}
