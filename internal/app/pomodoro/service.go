package pomodoro

import "github.com/ahmetcancicek/pomodorogo-server/internal/app/pomodoro/dto"

// StatisticService represent the statistic's service
type Service interface {
	Save(statDTO *dto.PomodoroDTO, userId uint) (*dto.PomodoroDTO, error)
}
