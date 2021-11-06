package application

import "context"

func (a App) DeleteCard(ctx context.Context, cardID, boardID string) error {
	return a.datastore.DeleteCard(ctx, cardID, boardID)
}

func (a App) DeleteColumn(ctx context.Context, columnID, boardID string) error {
	return a.datastore.DeleteColumn(ctx, columnID, boardID)
}
