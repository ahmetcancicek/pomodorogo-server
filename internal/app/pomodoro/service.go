package pomodoro

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

type Service interface {
	Save(pomodoro *model.Pomodoro, userId uint) (*model.Pomodoro, error)
}
