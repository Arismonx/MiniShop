package paymentRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Create interface and struct
type (
	PaymentRepositoryService interface{}

	paymentRepository struct {
		db *mongo.Client
	}
)

// Create constructor
func NewPaymentRepository(
	db *mongo.Client,
) PaymentRepositoryService {
	return &paymentRepository{db}
}

// call function for connect database
func (r *paymentRepository) paymentDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("payment_db")
}
