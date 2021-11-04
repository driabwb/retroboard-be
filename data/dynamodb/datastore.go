package dynamostore

type DynamoRetroboardDatastore struct {
	CardDynamoStore
	ColumnDynamoStore
	BoardDynamoStore
	EntityDynamoStore
}

func NewDynamoRetroboardDatastroe(client DynamoDB) DynamoRetroboardDatastore {
	return DynamoRetroboardDatastore{
		CardDynamoStore:   NewCardDynamoStore(client),
		ColumnDynamoStore: NewColumnDynamoStore(client),
		BoardDynamoStore:  NewBoardDynamoStore(client),
		EntityDynamoStore: NewEntityDynamoStore(client),
	}
}
