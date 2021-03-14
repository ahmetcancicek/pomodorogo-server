package service

import (
	"errors"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/statistic"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/statistic/dto"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/go-playground/validator/v10"
)

type statisticService struct {
	statisticRepository statistic.Repository
	tagRepository       tag.Repository
}

func NewStatisticService(statRepo statistic.Repository, tagRepo tag.Repository) statistic.Service {
	return &statisticService{
		statisticRepository: statRepo,
		tagRepository:       tagRepo,
	}
}

func (s statisticService) validate(statDTO *dto.StatisticDTO) error {
	err := utils.PayloadValidator(statDTO)
	if err != nil {
		return errors.New(err.(validator.ValidationErrors).Error())
	}
	return nil
}

func (s statisticService) Save(statDTO *dto.StatisticDTO, userId uint) (*dto.StatisticDTO, error) {
	// 1. Validate request data
	err := s.validate(statDTO)
	if err != nil {
		return nil, err
	}

	// 2. Checks if user have got his tag
	_, err = s.tagRepository.FindByIDAndUser(statDTO.TagID, userId)
	if err != nil {
		return nil, errors.New(model.ErrStatisticTagDoesNotExists)
	}

	// 3. Save
	stat := dto.ToStatistic(statDTO)
	stat, err = s.statisticRepository.Save(stat)

	return dto.ToStatisticDTO(stat), err
}
