package tag

import "github.com/ahmetcancicek/pomodorogo-server/internal/app/model"

// TagRepository represent the tag's repository
type Repository interface {
	FindByID(id int64) (*model.Tag, error)
	Save(tag *model.Tag) (*model.Tag, error)
	Update(tag *model.Tag) (*model.Tag, error)
	Delete(id int64) error
}
