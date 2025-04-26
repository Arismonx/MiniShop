package inventoryrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// *essential* Create interface and struct
type (
	InventoryRepositoryService interface{}

	inventoryRepository struct {
		db *mongo.Client
	}
)

// *essential* Create Constructor of inventoryRepository
func NewInventoryRepository(
	db *mongo.Client,
) InventoryRepositoryService {
	return &inventoryRepository{db}
}

// call function for connect database
func (r *inventoryRepository) inventoryDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("inventory_db")
}
