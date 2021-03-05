package service

import (
	"errors"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag/dto"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/go-playground/validator/v10"
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

func (t tagService) FindByID(id uint) (*dto.TagDTO, error) {
	tag, err := t.tagRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return dto.ToTagDTO(tag), err
}

func (t tagService) FindByName(name string) (*dto.TagDTO, error) {
	tag, err := t.tagRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return dto.ToTagDTO(tag), err
}

func (t tagService) FindByNameAndUser(name string, userId uint) (*dto.TagDTO, error) {
	tag, err := t.tagRepository.FindByNameAndUser(name, userId)
	if err != nil {
		return nil, err
	}

	return dto.ToTagDTO(tag), err
}

func (t tagService) FindByIDAndUser(id uint, userId uint) (*dto.TagDTO, error) {
	tag, err := t.tagRepository.FindByIDAndUser(id, userId)
	if err != nil {
		return nil, err
	}
	return dto.ToTagDTO(tag), err
}

func (t tagService) validate(tagDTO *dto.TagDTO) error {
	err := utils.PayloadValidator(tagDTO)
	if err != nil {
		return errors.New(err.(validator.ValidationErrors).Error())
	}
	return nil
}

func (t tagService) Save(tagDTO *dto.TagDTO, userId uint) (*dto.TagDTO, error) {

	err := t.validate(tagDTO)
	if err != nil {
		return nil, err
	}

	_, err = t.FindByNameAndUser(tagDTO.Name, userId)
	if err == nil {
		return nil, errors.New(model.ErrTagAlreadyExists)
	}

	tag := dto.ToTag(tagDTO)
	tag.UserID = userId
	tag, err = t.tagRepository.Save(tag)
	return dto.ToTagDTO(tag), err
}

func (t tagService) Update(tagDTO *dto.TagDTO) (*dto.TagDTO, error) {
	tag := dto.ToTag(tagDTO)
	tag.UpdatedAt = time.Now()
	tag, err := t.tagRepository.Update(tag)

	return dto.ToTagDTO(tag), err
}

func (t tagService) Delete(id uint) error {
	err := t.tagRepository.Delete(id)
	return err
}

func (t tagService) DeleteByIDAndUser(id uint, userId uint) error {
	err := t.tagRepository.DeleteByIDAndUser(id, userId)
	return err
}
