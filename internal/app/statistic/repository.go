package statistic

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// StatisticRepository represent the statistic's repository
type Repository interface {
	Save(statistic *model.Statistic) (*model.Statistic, error)
}
