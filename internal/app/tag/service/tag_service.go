package service

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag"
	"time"
)

type tagService struct {
	tagRepository tag.Repository
}

func NewTagService(tagRepository tag.Repository) tag.Service {
	return &tagService{
		tagRepository: tagRepository,
	}
}

func (t tagService) FindByID(id int64) (*model.Tag, error) {
	tag, err := t.tagRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (t tagService) Save(tag *model.Tag) error {

	// TODO: Name control
	err := t.tagRepository.Save(tag)
	if err != nil {
		return err
	}

	return nil
}

func (t tagService) Update(tag *model.Tag) error {
	tag.UpdatedAt = time.Now()
	return t.tagRepository.Update(tag)
}

func (t tagService) Delete(id int64) error {
	tag, err := t.tagRepository.FindByID(id)
	if err != nil {
		return err
	}

	if tag == nil {
		return model.ErrNotFound
	}

	return t.tagRepository.Delete(id)
}
