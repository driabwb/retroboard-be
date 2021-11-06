package application

import "context"

func (a App) UpdateBoardTitle(ctx context.Context, boardID, title string) error {
	return a.datastore.UpdateBoardTitle(ctx, boardID, title)
}
