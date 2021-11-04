package datastore

type RetroboardDatastore interface {
	CardStorer
	CardColumnUpdater
	CardContentUpdater
	CardVoteUpdater
	CardDeleter
	ColumnStorer
	ColumnTitleUpdater
	ColumnDeleter
	BoardStorer
	BoardTitleUpdater
	EntityFetcher
}
