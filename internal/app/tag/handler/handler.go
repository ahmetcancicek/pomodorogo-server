package handler

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/handler"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag/dto"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type TagHandler struct {
	logger     *logrus.Logger
	TagService tag.Service
}

func NewTagHandler(r *mux.Router, log *logrus.Logger, ts tag.Service, mf mux.MiddlewareFunc) {
	tagHandler := &TagHandler{
		logger:     log,
		TagService: ts,
	}

	r.HandleFunc("/api/v1/tags", tagHandler.create).Methods(http.MethodPost)
	r.Use(mf)
}

func (h TagHandler) create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tagDTO := new(dto.TagDTO)

	// 1. Decode request body
	err := utils.FromJSON(tagDTO, r.Body)
	if err != nil {
		h.logger.Error("unable to decode user json", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.Error()}, w)
		return
	}
	defer r.Body.Close()

	// 2. Validate
	err = utils.PayloadValidator(tagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.(validator.ValidationErrors).Error()}, w)
		return
	}

	// 3. Check if tag exist in database
	_, err = h.TagService.FindByName(tagDTO.Name)
	if err == nil {
		utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: model.ErrTagAlreadyExists}, w)
		return
	}

	userID := r.Context().Value(handler.UserIDKey{}).(int64)

	tag := new(model.Tag)
	tag.UserID = userID
	tag.Name = tagDTO.Name
	tag.Colour = tagDTO.Colour
	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()
	tag, err = h.TagService.Save(tag)
	if err != nil {
		h.logger.Error("unable to insert tag to database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrTagCreateFailed}, w)
		return
	}

	// 6- Respond successful message
	h.logger.Debug("tag created successfully")
	w.WriteHeader(http.StatusCreated)
	utils.ToJSON(&model.GenericResponse{Code: 200, Status: true, Message: "Tag created successfully", Data: tag}, w)

}
