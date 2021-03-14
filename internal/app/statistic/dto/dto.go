package dto

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"time"
)

// StatisticDTO ...
type StatisticDTO struct {
	ID         uint      `json:"id"`
	StartTime  time.Time `json:"startTime"`
	FinishTime time.Time `json:"finishTime"`
	TagID      uint      `json:"tagId"`
}

// ToStatistic ...
func ToStatistic(statDTO *StatisticDTO) *model.Statistic {
	return &model.Statistic{
		ID:         statDTO.ID,
		StartTime:  statDTO.StartTime,
		FinishTime: statDTO.FinishTime,
		TagID:      statDTO.TagID,
	}
}

// ToStatisticDTO
func ToStatisticDTO(statistic *model.Statistic) *StatisticDTO {
	return &StatisticDTO{
		ID:         statistic.ID,
		StartTime:  statistic.StartTime,
		FinishTime: statistic.FinishTime,
		TagID:      statistic.TagID,
	}
}
