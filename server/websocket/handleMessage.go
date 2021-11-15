package websocket

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/driabwb/retroboard/domain"
	"github.com/driabwb/retroboard/messages"
)

func (c *client) handleMessage(msg messages.Message) (isClosing bool) {
	// TODO: What else to put on the context?
	ctx := context.Background()

	switch msg.Type {
	case messages.MessageCreateCard:
		payload := messages.MessageCreateCardPayload{}
		err := json.Unmarshal(msg.Msg, &payload)
		if err != nil {
			// Do Something
			c.Send(messages.NewMessageError(
				msg,
				fmt.Errorf("Failed to unmarshal message: %w", err),
			))
			break
		}
		card := domain.Card{
			ColumnID: payload.ColumnID,
			BoardID:  payload.BoardID,
			Content:  payload.Content,
			Votes:    0,
		}
		newCard, err = c.app.CreateCard(ctx, card)
		if err != nil {
			c.Send(messages.NewMessageError(msg, err))
			break
		}

		resp := messages.Message{
			Type: msg.Type,
			Msg: messages.MessageCreateCardPayload{
				CardID:   newCard.ID,
				ColumnID: newCard.ColumnID,
				BoardID:  newCard.BoardID,
				Content:  newCard.Content,
			},
		}

		c.Send(resp)
		c.pool.Broadcast(c.id, resp)
	case messages.MessageUpdateCardContent:
		payload := messages.MessageUpdateCardContentPayload{}
		err := json.Unmarshal(msg.Msg, &payload)
		if err != nil {
			c.Send(messages.NewMessageError(
				msg,
				fmt.Errorf("Failed to unmarshal message: %w", err),
			))
			break
		}

		err = c.app.UpdateCardContent(ctx, payload.CardID, payload.BoardID, payload.Content)
		if err != nil {
			c.Send(messages.NewMessageError(msg, err))
		}
		c.pool.Broadcast(c.id, msg)
	case messages.MessageUpdateCardColumn:
		payload := messages.MessageUpdateCardColumnPayload{}
		err := json.Unmarshal(msg.Msg, &payload)
		if err != nil {
			c.Send(messages.NewMessageError(
				msg,
				fmt.Errorf("Failed to unmarshal message: %w", err),
			))
			break
		}

		err = c.app.UpdateCardColumn(ctx, payload.CardID, payload.ColumnID, card.BoardID)
		if err != nil {
			c.Send(messages.NewMessageError(msg, err))
		}
		c.pool.Broadcast(c.id, msg)
	case messages.MessageUpdateCardVotes:
		payload := messages.MessageUpdateCardContentPayload{}
		err := json.Unmarshal(msg.Msg, &payload)
		if err != nil {
			c.Send(messages.NewMessageError(
				msg,
				fmt.Errorf("Failed to unmarshal message: %w", err),
			))
			break
		}

		err = c.app.UpdateCardVotes(ctx, payload.CardID, payload.BoardID, payload.Delta)
		if err != nil {
			c.Send(messages.NewMessageError(msg, err))
		}
		c.pool.Broadcast(c.id, msg)
	case messages.MessageDeleteCard:
		payload := messages.MessageUpdateCardContentPayload{}
		err := json.Unmarshal(msg.Msg, &payload)
		if err != nil {
			c.Send(messages.NewMessageError(
				msg,
				fmt.Errorf("Failed to unmarshal message: %w", err),
			))
			break
		}

		err = c.app.DeleteCard(ctx, payload.CardID, payload.BoardID)
		if err != nil {
			c.Send(messages.NewMessageError(msg, err))
		}
		c.pool.Broadcast(c.id, msg)
	case messages.MessageCreateColumn:
		payload := messages.MessageUpdateCardContentPayload{}
		err := json.Unmarshal(msg.Msg, &payload)
		if err != nil {
			c.Send(messages.NewMessageError(
				msg,
				fmt.Errorf("Failed to unmarshal message: %w", err),
			))
			break
		}

		col := domain.Column{
			BoardID: payload.BoardID,
			Title:   payload.Title,
		}
		newCol, err := c.app.CreateColumn(ctx, col)
		if err != nil {
			c.Send(messages.NewMessageError(msg, err))
			break
		}

		resp := messages.Message{
			Type: msg.Type,
			Msg: messages.MessageCreateColumn{
				ColumnID: newCol.ID,
				BoardID:  newCol.BoardID,
				Title:    newCol.Title,
			},
		}

		c.Send(resp)
		c.pool.Broadcast(c.id, resp)
	case messages.MessageUpdateColumnTitle:
		payload := messages.MessageUpdateCardContentPayload{}
		err := json.Unmarshal(msg.Msg, &payload)
		if err != nil {
			c.Send(messages.NewMessageError(
				msg,
				fmt.Errorf("Failed to unmarshal message: %w", err),
			))
			break
		}

		err = c.app.UpdateColumnTitle(ctx, payload.ColumnID, payload.BoardID, payload.Title)
		if err != nil {
			c.Send(messages.NewMessageError(msg, err))
		}
		c.pool.Broadcast(c.id, msg)
	case messages.MessageDeleteColumn:
		payload := messages.MessageUpdateCardContentPayload{}
		err := json.Unmarshal(msg.Msg, &payload)
		if err != nil {
			c.Send(messages.NewMessageError(
				msg,
				fmt.Errorf("Failed to unmarshal message: %w", err),
			))
			break
		}

		err = c.app.DeleteColumn(ctx, payload.ColumnID, payload.BoardID)
		if err != nil {
			c.Send(messages.NewMessageError(msg, err))
		}
		c.pool.Broadcast(c.id, msg)
	case messages.MessageUpdateBoardTitle:
		payload := messages.MessageUpdateCardContentPayload{}
		err := json.Unmarshal(msg.Msg, &payload)
		if err != nil {
			c.Send(messages.NewMessageError(
				msg,
				fmt.Errorf("Failed to unmarshal message: %w", err),
			))
			break
		}

		err = c.app.UpdateBoardTitle(ctx, payload.BoardID, payload.Title)
		if err != nil {
			c.Send(messages.NewMessageError(msg, err))
		}
		c.pool.Broadcast(c.id, msg)
	default:
		c.Send(messages.NewMessageError(
			msg,
			messages.ErrorUnknownMessageType(msg.Type),
		))
	}
}
