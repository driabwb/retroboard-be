package datastore

import "context"

type EntityFetcher interface {
	GetEntity(ctx context.Context, boardID string) (Entity, error)
}
