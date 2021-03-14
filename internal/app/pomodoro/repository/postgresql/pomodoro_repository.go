package postgresql

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/pomodoro"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type postgreSQLPomodoroRepository struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewPostgreSQLPomodoroRepository(log *logrus.Logger, db *gorm.DB) pomodoro.Repository {
	return &postgreSQLPomodoroRepository{
		logger: log,
		db:     db,
	}
}

func (p postgreSQLPomodoroRepository) Save(pomodoro *model.Pomodoro) (*model.Pomodoro, error) {
	p.logger.Info("creating pomodoro:", pomodoro)
	err := p.db.Create(&pomodoro).Error
	return pomodoro, err
}

func CreateRepository(db *gorm.DB) pomodoro.Repository {
	return &postgreSQLPomodoroRepository{
		logger: utils.NewLogger(),
		db:     db,
	}
}
