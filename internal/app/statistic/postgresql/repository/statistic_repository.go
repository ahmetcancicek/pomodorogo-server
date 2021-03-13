package postgresql

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/statistic"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type postgreSQLStatisticRepository struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewPostgreSQLStatisticRepository(l *logrus.Logger, db *gorm.DB) statistic.Repository {
	return &postgreSQLStatisticRepository{
		logger: l,
		db:     db,
	}
}

func (p postgreSQLStatisticRepository) Save(statistic *model.Statistic) (*model.Statistic, error) {
	p.logger.Info("creating statistic:", statistic)
	err := p.db.Create(&statistic).Error
	return statistic, err
}

func CreateRepository(db *gorm.DB) statistic.Repository {
	return &postgreSQLStatisticRepository{
		logger: utils.NewLogger(),
		db:     db,
	}
}
