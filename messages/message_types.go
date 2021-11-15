package messages

import "fmt"

type Message struct {
	Type string `json:"type"`
	Msg  []byte `json:"message"`
}

const (
	MessageCreateCard        = "MessageCreateCard"
	MessageUpdateCardContent = "MessageUpdateCardContent"
	MessageUpdateCardColumn  = "MessageUpdateCardColumn"
	MessageUpdateCardVotes   = "MessageUpdateCardVotes"
	MessageDeleteCard        = "MessageDeleteCard"

	MessageCreateColumn      = "MessageCreateColumn"
	MessageUpdateColumnTitle = "MessageUpdateColumnTitle"
	MessageDeleteColumn      = "MessageDeleteColumn"

	MessageUpdateBoardTitle = "MessageUpdateBoardTitle"
)

var (
	ErrorUnknownMessageType = func(msgType string) error { return fmt.Errorf("Unknown Messag Type: %s", msgType) }
)

type (
	MessageCreateCardPayload struct {
		CardID   string `json:"card_id"`
		ColumnID string `json:"column_id"`
		BoardID  string `json:"board_id"`
		Content  string `json:"content"`
	}
	MessageUpdateCardContentPayload struct {
		CardID  string `json:"card_id"`
		BoardID string `json:"board_id"`
		Content string `json:"content"`
	}
	MessageUpdateCardColumnPayload struct {
		CardID   string `json:"card_id"`
		ColumnID string `json:"column_id"`
		BoardID  string `json:"board_id"`
	}
	MessageUpdateCardVotesPayload struct {
		CardID  string `json:"card_id"`
		BoardID string `json:"board_id"`
		Delta   int    `json:"delta"`
	}
	MessageDeleteCardPayload struct {
		CardID  string `json:"card_id"`
		BoardID string `json:"board_id"`
	}

	MessageCreateColumnPayload struct {
		ColumnID string `json:"column_id"`
		BoardID  string `json:"board_id"`
		Title    string `json:"title"`
	}
	MessageUpdateColumnTitlePayload struct {
		ColumnID string `json:"column_id"`
		BoardID  string `json:"board_id"`
		Title    string `json:"title"`
	}
	MessageDeleteColumnPayload struct {
		ColumnID string `json:"column_id"`
		BoardID  string `json:"board_id"`
	}

	MessageUpdateBoardTitlePayload struct {
		BoardID string `json:"board_id"`
		Title   string `json:"title"`
	}

	MessageErrorPayload struct {
		MessageType string `json:"message_type"`
		Error       error  `json:"error"`
	}
)

func NewMessageError(request Message, err error) Message {
	return Message{
		Type: messages.MessageError,
		Msg: messages.MessageErrorPayload{
			MessageType: msg.Type,
			Error:       err,
		},
	}
}
