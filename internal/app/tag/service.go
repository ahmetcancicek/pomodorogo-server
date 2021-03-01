package tag

import "github.com/ahmetcancicek/pomodorogo-server/internal/app/model"

// TagService represent the tag's service
type Service interface {
	FindByID(id int64) (*model.Tag, error)
	Save(tag *model.Tag) error
	Update(tag *model.Tag) error
	Delete(id int64) error
}
