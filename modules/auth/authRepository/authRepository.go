package authRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// *essential* Create interface and struct
type (
	AuthRepositoryService interface{}

	authRepository struct {
		db *mongo.Client
	}
)

// *essential* Create Constructor of struct authRepository
func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db}
}

// call function for connect database
func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}
