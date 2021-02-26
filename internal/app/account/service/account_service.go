package service

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/account"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"time"
)

type accountService struct {
	accountRepository account.Repository
}

// NewAccountService will create new an accountService object representation of of account.Service interface
func NewAccountService(accountRepository account.Repository) account.Service {
	return &accountService{
		accountRepository: accountRepository,
	}
}

func (u accountService) FindByID(id int64) (*model.User, error) {
	user, err := u.accountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u accountService) FindByUUID(uuid string) (*model.User, error) {
	user, err := u.accountRepository.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u accountService) FindByEmail(email string) (*model.User, error) {
	user, err := u.accountRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u accountService) Update(user *model.User) error {
	user.UpdatedAt = time.Now()
	return u.accountRepository.Update(user)
}

func (u accountService) Save(user *model.User) error {

	// TODO: Username, email control
	err := u.accountRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (u accountService) Delete(id int64) error {
	user, err := u.accountRepository.FindByID(id)
	if err != nil {
		return err
	}

	if user == nil {
		return model.ErrNotFound
	}

	return u.accountRepository.Delete(id)
}
