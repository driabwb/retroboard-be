package application

import "context"

func (a App) UpdateCardContent(ctx context.Context, cardID, boardID, content string) error {
	return a.datastore.UpdateCardContent(ctx, cardID, boardID, content)
}

func (a App) UpdateCardColumn(ctx context.Context, cardID, columnID, boardID string) error {
	return a.datastore.UpdateCardContent(ctx, cardID, columnID, boardID)
}

func (a App) UpdateCardVotes(ctx context.Context, cardID, boardID string, delta int) error {
	return a.datastore.AddCardVotes(ctx, cardID, boardID, delta)
}
