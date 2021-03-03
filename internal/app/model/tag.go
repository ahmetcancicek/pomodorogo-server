package model

import (
	"time"
)

// Tag ...
type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Colour    string    `json:"colour"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    int64
}
