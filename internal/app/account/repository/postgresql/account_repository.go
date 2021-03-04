package postgresql

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/account"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type postgreSQLAccountRepository struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewPostgreSQLAccountRepository(log *logrus.Logger, db *gorm.DB) account.Repository {
	return &postgreSQLAccountRepository{
		logger: log,
		db:     db,
	}
}

func (p postgreSQLAccountRepository) FindByID(id int64) (*model.User, error) {
	p.logger.Debug("finding for user with id", id)
	user := new(model.User)
	err := p.db.Where(`id = ?`, id).First(&user).Error
	p.logger.Debug("read user", user)
	return user, err
}

func (p postgreSQLAccountRepository) FindByUUID(uuid string) (*model.User, error) {
	p.logger.Debug("finding for user with uuid", uuid)
	user := new(model.User)
	err := p.db.Where(`uuid = ?`, uuid).First(&user).Error
	p.logger.Debug("read user", user)
	return user, err
}

func (p postgreSQLAccountRepository) FindByEmail(email string) (*model.User, error) {
	p.logger.Debug("finding for user with email", email)
	user := new(model.User)
	err := p.db.Where(`email = ?`, email).First(&user).Error
	p.logger.Debug("read user", user)
	return user, err
}

func (p postgreSQLAccountRepository) FindByUsername(username string) (*model.User, error) {
	p.logger.Debug("finding for user with username", username)
	user := new(model.User)
	err := p.db.Where(`username = ?`, username).First(&user).Error
	p.logger.Debug("read user", user)
	return user, err
}

func (p postgreSQLAccountRepository) Update(user *model.User) error {
	// TODO:
	return nil
}

func (p postgreSQLAccountRepository) Save(user *model.User) error {
	p.logger.Info("creating user", user)
	err := p.db.Create(&user).Error
	return err
}

func (p postgreSQLAccountRepository) Delete(id int64) error {
	p.logger.Info("deleting user with id", id)
	err := p.db.Delete(&model.User{}, id).Error
	return err
}

func CreateRepository(db *gorm.DB) account.Repository {
	return &postgreSQLAccountRepository{
		logger: utils.NewLogger(),
		db:     db,
	}
}
