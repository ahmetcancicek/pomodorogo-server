package account

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// AccountRepository represent the account's repository
type Repository interface {
	FindByID(id int64) (*model.User, error)
	FindByUUID(uuid string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Save(user *model.User) error
	Delete(id int64) error
}
