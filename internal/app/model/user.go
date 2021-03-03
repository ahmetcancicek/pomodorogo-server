package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// User ...
type User struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid; type:varchar(100)`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	TokenHash string    `json:"tokenHash"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
