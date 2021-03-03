package dto

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// Tag ...
type TagDTO struct {
	Name   string `json:"name" validate:"max=50"`
	Colour string `json:"colour" validate:"max=7"`
}

// ToTag ...
func ToTag(tagDTO *TagDTO) *model.Tag {
	return &model.Tag{
		Name:   tagDTO.Name,
		Colour: tagDTO.Colour,
	}
}

// ToTagDTO ...
func ToTagDTO(tag *model.Tag) *TagDTO {
	return &TagDTO{
		Name:   tag.Name,
		Colour: tag.Colour,
	}
}
