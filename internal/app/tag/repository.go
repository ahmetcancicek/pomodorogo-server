package tag

import "github.com/ahmetcancicek/pomodorogo-server/internal/app/model"

// TagRepository represent the tag's repository
type Repository interface {
	FindByID(id int64) (*model.Tag, error)
	Save(label *model.Tag) error
	Update(label *model.Tag) error
	Delete(id int64) error
}
