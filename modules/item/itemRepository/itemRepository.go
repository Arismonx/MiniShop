package itemRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Create interface and struct
type (
	ItemRepositoryService interface{}

	itemRepository struct {
		db *mongo.Client
	}
)

// Create constructor
func NewItemRepository(db *mongo.Client) ItemRepositoryService {
	return &itemRepository{db}
}

// call function for connect database
func (r *itemRepository) itemDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("item_db")
}
