package account

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// AccountService represent the account's service
type Service interface {
	FindByID(id int64) (*model.User, error)
	FindByUUID(uuid string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Save(user *model.User) error
	Delete(id int64) error
}
