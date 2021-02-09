package auth

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// UserRepository represent the user's repository
type Repository interface {
	FindByID(id int64) (*model.User, error)
	Update(user *model.User) error
	Save(user *model.User) error
	Delete(id int64) error
}
