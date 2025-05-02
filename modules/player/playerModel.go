package player

import "time"

type (
	PlayerProfile struct {
		Id         string    `json:"id"`
		Email      string    `json:"email"`
		Username   string    `json:"username"`
		Created_at time.Time `json:"created_at"`
		Updated_at time.Time `json:"updated_at"`
	}
)
