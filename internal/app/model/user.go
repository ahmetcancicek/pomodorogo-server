package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// User ...
type User struct {
	ID        int64     `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	TokenHash string    `json:"tokenHash"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
