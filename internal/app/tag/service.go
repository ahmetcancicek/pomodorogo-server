package tag

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag/dto"
)

// TagService represent the tag's service
type Service interface {
	FindByID(id uint) (*dto.TagDTO, error)
	FindByName(name string) (*dto.TagDTO, error)
	Save(tagDTO *dto.TagDTO, userId uint) (*dto.TagDTO, error)
	Update(tagDTO *dto.TagDTO) (*dto.TagDTO, error)
	Delete(id uint) error
}
