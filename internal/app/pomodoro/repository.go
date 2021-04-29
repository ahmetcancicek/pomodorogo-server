package pomodoro

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

type Repository interface {
	Save(pomodoro *model.Pomodoro) (*model.Pomodoro, error)
}
