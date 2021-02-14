package user

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// UserService represent the user's service
type Service interface {
	FindByID(id int64) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByCredentials(email, password string) (*model.User, error)
	Update(user *model.User) error
	Save(user *model.User) error
	Delete(id int64) error
}
