package datastore

import (
	"context"

	"github.com/driabwb/retroboard/domain"
)

type BoardStorer interface {
	SaveBoard(ctx context.Context, board domain.Board) error
}

type BoardTitleUpdater interface {
	UpdateBoardTitle(ctx context.Context, id, title string) error
}
