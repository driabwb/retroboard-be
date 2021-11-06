package application

import "context"

func (a App) UpdateColumnTitle(ctx context.Context, columnID, boardID, title string) error {
	return a.datastore.UpdateColumnTitle(ctx, columnID, boardID, title)
}
