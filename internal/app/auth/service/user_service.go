package service

import (
	"context"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"time"
)

type userService struct {
	userRepository auth.Repository
	contextTimeout time.Duration
}

// NewUserService will create new an useService object representation of of user.Service interface
func NewUserService(userRepository auth.Repository, timeout time.Duration) auth.Service {
	return &userService{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u userService) FindByID(ctx context.Context, id int64) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userService) Update(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user.UpdatedAt = time.Now()
	return u.userRepository.Update(ctx, user)
}

func (u userService) Save(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	// TODO: Username, email control
	err := u.userRepository.Save(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if user == nil {
		return model.ErrNotFound
	}

	return u.userRepository.Delete(ctx, id)
}
