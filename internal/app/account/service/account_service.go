package service

import (
	"errors"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/account"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
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

func (u accountService) validate(user *model.User) error {
	err := utils.PayloadValidator(user)
	if err != nil {
		return errors.New(err.(validator.ValidationErrors).Error())
	}
	return nil
}

func (u accountService) Save(user *model.User) error {

	err := u.validate(user)
	if err != nil {
		return err
	}

	_, err = u.FindByEmail(user.Email)
	if err == nil {
		return errors.New(model.ErrTagAlreadyExists)
	}

	// TODO: Username control
	user.UUID = uuid.NewV4()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = u.accountRepository.Save(user)
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
