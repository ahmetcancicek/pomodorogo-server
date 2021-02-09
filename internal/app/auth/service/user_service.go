package service

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"time"
)

type userService struct {
	userRepository auth.Repository
}

// NewUserService will create new an useService object representation of of user.Service interface
func NewUserService(userRepository auth.Repository) auth.Service {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) FindByID(id int64) (*model.User, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userService) Update(user *model.User) error {
	user.UpdatedAt = time.Now()
	return u.userRepository.Update(user)
}

func (u userService) Save(user *model.User) error {

	// TODO: Username, email control
	err := u.userRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) Delete(id int64) error {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return err
	}

	if user == nil {
		return model.ErrNotFound
	}

	return u.userRepository.Delete(id)
}
