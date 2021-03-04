package postgresql

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type postgreSQLTagRepository struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewPostgreSQLTagRepository(l *logrus.Logger, db *gorm.DB) tag.Repository {
	return &postgreSQLTagRepository{
		logger: l,
		db:     db,
	}
}

func (p postgreSQLTagRepository) FindByID(id uint) (*model.Tag, error) {
	p.logger.Debug("finding for tag with id", id)
	label := new(model.Tag)
	err := p.db.Where(`id = ?`, id).First(&label).Error
	p.logger.Debug("read tag", label)
	return label, err

}

func (p postgreSQLTagRepository) FindByName(name string) (*model.Tag, error) {
	p.logger.Debug("finding for tag with name", name)
	label := new(model.Tag)
	err := p.db.Where(`name = ?`, name).First(&label).Error
	p.logger.Debug("read tag", label)
	return label, err
}

func (p postgreSQLTagRepository) Save(label *model.Tag) (*model.Tag, error) {
	p.logger.Info("creating tag", label)
	err := p.db.Create(&label).Error
	return label, err
}

func (p postgreSQLTagRepository) Update(label *model.Tag) (*model.Tag, error) {
	// TODO:
	return label, nil
}

func (p postgreSQLTagRepository) Delete(id uint) error {
	p.logger.Info("deleting tag with id", id)
	err := p.db.Delete(&model.Tag{}, id).Error
	return err
}

func CreateRepository(db *gorm.DB) tag.Repository {
	return &postgreSQLTagRepository{
		logger: utils.NewLogger(),
		db:     db,
	}
}
