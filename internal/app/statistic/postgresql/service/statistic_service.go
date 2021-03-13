package service

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/statistic"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/statistic/dto"
)

type statisticService struct {
	statisticRepository statistic.Repository
}

func NewStatisticService(sr statistic.Repository) statistic.Service {
	return &statisticService{
		statisticRepository: sr,
	}
}

func (s statisticService) Save(statDTO *dto.StatisticDTO) (*dto.StatisticDTO, error) {
	// TODO:
	// Should validate request data
	// Checks if user have got this label
	//

	stat := dto.ToStatistic(statDTO)
	stat, err := s.statisticRepository.Save(stat)

	return dto.ToStatisticDTO(stat), err
}
