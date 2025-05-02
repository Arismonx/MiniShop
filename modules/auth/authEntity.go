package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Credential struct {
		Id           primitive.ObjectID `json:"_id" bson:"_id",omitempty`
		PlayerId     string             `json:"player_id" bson:"player_id"`
		RoleCode     int                `json:"role_code" bson:"role_code"`
		AccessToken  string             `json:"access_token" bson:"access_token"`
		RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
		Created_at   time.Time          `json:"created_at"`
		Updated_at   time.Time          `json:"updated_at"`
	}

	Role struct {
		Id    primitive.ObjectID `json:"_id" bson:"_id",omitempty`
		Titel string             `json:"tile" bson:"title"`
		Code  string             `json:"code" bson:"code"`
	}
)
