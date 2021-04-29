package service

import (
	"errors"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/pomodoro"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/go-playground/validator/v10"
)

type pomodoroService struct {
	pomodoroRepository pomodoro.Repository
	tagRepository      tag.Repository
}

func NewStatisticService(pomodoroRepository pomodoro.Repository, tagRepo tag.Repository) pomodoro.Service {
	return &pomodoroService{
		pomodoroRepository: pomodoroRepository,
		tagRepository:      tagRepo,
	}
}

func (s pomodoroService) validate(pomodoro *model.Pomodoro) error {
	err := utils.PayloadValidator(pomodoro)
	if err != nil {
		return errors.New(err.(validator.ValidationErrors).Error())
	}
	return nil
}

func (s pomodoroService) Save(pomodoro *model.Pomodoro, userId uint) (*model.Pomodoro, error) {
	// 1. Validate request data
	err := s.validate(pomodoro)
	if err != nil {
		return nil, err
	}

	// 2. Checks if user have got his tag
	_, err = s.tagRepository.FindByIDAndUser(pomodoro.TagID, userId)
	if err != nil {
		return nil, errors.New(model.ErrPomodoroTagDoesNotExists)
	}

	// 3. Save
	return s.pomodoroRepository.Save(pomodoro)
}
