package tag

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// TagService represent the tag's service
type Service interface {
	FindByID(id uint) (*model.Tag, error)
	FindByName(name string) (*model.Tag, error)
	FindByNameAndUser(name string, userId uint) (*model.Tag, error)
	FindByIDAndUser(id uint, userId uint) (*model.Tag, error)
	Save(tag *model.Tag, userId uint) (*model.Tag, error)
	Update(tag *model.Tag) (*model.Tag, error)
	Delete(id uint) error
	DeleteByIDAndUser(id uint, userId uint) error
}
