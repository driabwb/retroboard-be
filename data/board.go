package datastore

import (
	"context"

	"github.com/driabwb/retroboard/domain"
)

type BoardStorer interface {
	SaveBoard(ctx context.Context, board domain.Board) error
}

type BoardTitleUpdater interface {
	UpdateTitle(ctx context.Context, id, title string) error
}
