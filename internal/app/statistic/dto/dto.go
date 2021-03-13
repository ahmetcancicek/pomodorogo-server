package dto

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"time"
)

// StatisticDTO ...
type StatisticDTO struct {
	ID         uint      `json:"id"`
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
	TagID      uint
}

// ToStatistic ...
func ToStatistic(statDTO *StatisticDTO) *model.Statistic {
	return &model.Statistic{
		ID:         statDTO.ID,
		StartedAt:  statDTO.StartedAt,
		FinishedAt: statDTO.FinishedAt,
	}
}

// ToStatisticDTO
func ToStatisticDTO(statistic *model.Statistic) *StatisticDTO {
	return &StatisticDTO{
		ID:         statistic.ID,
		StartedAt:  statistic.StartedAt,
		FinishedAt: statistic.FinishedAt,
	}
}
