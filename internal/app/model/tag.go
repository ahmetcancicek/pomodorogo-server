package model

import (
	"time"
)

// Tag ...
type Tag struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	Name      string     `json:"name"`
	Colour    string     `json:"colour"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
