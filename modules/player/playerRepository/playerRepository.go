package playerRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Create interface and struct
type (
	PlayerRepositoryService interface{}

	playerRepository struct {
		db *mongo.Client
	}
)

// Create constructor
func NewPlayerRepository(
	db *mongo.Client,
) PlayerRepositoryService {
	return &playerRepository{db}
}

// call function for connect database
func (r *playerRepository) PlayerDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}
