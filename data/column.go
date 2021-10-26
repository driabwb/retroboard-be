package datastore

import (
	"context"

	"github.com/driabwb/retroboard/domain"
)

type ColumnStorer interface {
	SaveColumn(ctx context.Context, col domain.Column) error
}

type ColumnDeleter interface {
	DeleteColumn(ctx context.Context, id, boardID string) error
}

type ColumnUpdater interface {
	UpdateTitle(ctx context.Context, id, boardID, title string) error
}
