package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// function convert string to BSON
func ConvertToObject(id string) primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return objectID
}
