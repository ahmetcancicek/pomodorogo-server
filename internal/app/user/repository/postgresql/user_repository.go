package postgresql

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/user"
	"gorm.io/gorm"
)

type postgreSQLUserRepository struct {
	db *gorm.DB
}

func NewPostgreSQLUserRepository(db *gorm.DB) user.Repository {
	return &postgreSQLUserRepository{db}
}

func (p postgreSQLUserRepository) FindByID(id int64) (*model.User, error) {
	user := new(model.User)
	err := p.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

func (p postgreSQLUserRepository) FindByUUID(uuid string) (*model.User, error) {
	user := new(model.User)
	err := p.db.Where(`uuid = ?`, uuid).First(&user).Error
	return user, err
}

func (p postgreSQLUserRepository) FindByEmail(email string) (*model.User, error) {
	user := new(model.User)
	err := p.db.Where(`email = ?`, email).First(&user).Error
	return user, err
}

func (p postgreSQLUserRepository) Update(user *model.User) error {
	// TODO:
	return nil
}

func (p postgreSQLUserRepository) Save(user *model.User) error {
	err := p.db.Create(&user).Error
	return err
}

func (p postgreSQLUserRepository) Delete(id int64) error {
	err := p.db.Delete(&model.User{}, id).Error
	return err
}

func CreateRepository(db *gorm.DB) user.Repository {
	return &postgreSQLUserRepository{db: db}
}
