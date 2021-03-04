package account

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// AccountService represent the account's service
type Service interface {
	FindByID(id uint) (*model.User, error)
	FindByUUID(uuid string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Update(user *model.User) error
	Save(user *model.User) error
	Delete(id uint) error
}
