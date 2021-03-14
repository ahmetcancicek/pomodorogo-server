package statistic

import "github.com/ahmetcancicek/pomodorogo-server/internal/app/statistic/dto"

// StatisticService represent the statistic's service
type Service interface {
	Save(statDTO *dto.StatisticDTO, userId uint) (*dto.StatisticDTO, error)
}
