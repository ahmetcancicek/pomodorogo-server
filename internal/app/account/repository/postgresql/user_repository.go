package postgresql

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/account"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"gorm.io/gorm"
)

type postgreSQLAccountRepository struct {
	db *gorm.DB
}

func NewPostgreSQLAccountRepository(db *gorm.DB) account.Repository {
	return &postgreSQLAccountRepository{db}
}

func (p postgreSQLAccountRepository) FindByID(id int64) (*model.User, error) {
	user := new(model.User)
	err := p.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

func (p postgreSQLAccountRepository) FindByUUID(uuid string) (*model.User, error) {
	user := new(model.User)
	err := p.db.Where(`uuid = ?`, uuid).First(&user).Error
	return user, err
}

func (p postgreSQLAccountRepository) FindByEmail(email string) (*model.User, error) {
	user := new(model.User)
	err := p.db.Where(`email = ?`, email).First(&user).Error
	return user, err
}

func (p postgreSQLAccountRepository) Update(user *model.User) error {
	// TODO:
	return nil
}

func (p postgreSQLAccountRepository) Save(user *model.User) error {
	err := p.db.Create(&user).Error
	return err
}

func (p postgreSQLAccountRepository) Delete(id int64) error {
	err := p.db.Delete(&model.User{}, id).Error
	return err
}

func CreateRepository(db *gorm.DB) account.Repository {
	return &postgreSQLAccountRepository{db: db}
}
