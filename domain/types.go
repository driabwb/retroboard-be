package domain

type (
	Card struct {
		ID       string
		ColumnID string
		BoardID  string
		Content  string
		Votes    uint
	}

	Column struct {
		ID      string
		BoardID string
		Title   string
	}

	Board struct {
		ID    string
		Title string
	}

	Entity struct {
		Board
		Cols  []Column
		Cards []Card
	}
)
