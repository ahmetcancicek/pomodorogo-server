package handler

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/handler"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag/dto"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	r.HandleFunc("/api/v1/tags/{id:[0-9]+}", tagHandler.read).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/tags", tagHandler.update).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/tags/{id:[0-9]+}", tagHandler.delete).Methods(http.MethodDelete)
	r.Use(mf)
}

func (h TagHandler) create(w http.ResponseWriter, r *http.Request) {

	tagDTO := new(dto.TagDTO)

	// 1. Decode request body
	err := utils.FromJSON(tagDTO, r.Body)
	if err != nil {
		h.logger.Error("unable to decode tag json", err.Error())
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.Error()})
		return
	}
	defer r.Body.Close()

	// 2. Save
	userId := r.Context().Value(handler.UserIDKey{}).(uint)
	createdTag, err := h.TagService.Save(dto.ToTag(tagDTO), userId)
	if err != nil {
		h.logger.Error("unable to insert tag to database: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: err.Error()})
		return
	}

	// 3. Respond successful message
	h.logger.Debug("tag created successfully")
	utils.RespondWithJSON(w,
		&model.GenericResponse{Code: http.StatusCreated, Status: true, Message: "Tag created successfully", Data: dto.ToTagDTO(createdTag)})

}

func (h TagHandler) read(w http.ResponseWriter, r *http.Request) {
	// Integer control
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("unable to get parameter because of variable type: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.Error()})
		return
	}

	userId := r.Context().Value(handler.UserIDKey{}).(uint)

	// 2. Get
	tag, err := h.TagService.FindByIDAndUser(uint(id), userId)
	if err != nil {
		h.logger.Error("unable to get tag to database: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: err.Error()})
		return
	}

	// 3- Respond successful message
	h.logger.Debug("tag got successfully")
	utils.RespondWithJSON(w,
		&model.GenericResponse{Code: http.StatusOK, Status: true, Message: "Tag got successfully", Data: dto.ToTagDTO(tag)})
}

func (h TagHandler) update(w http.ResponseWriter, r *http.Request) {

}

func (h TagHandler) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Integer control
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("unable to get parameter because of variable type: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.Error()})
		return
	}

	userId := r.Context().Value(handler.UserIDKey{}).(uint)

	// 2. Delete
	err = h.TagService.DeleteByIDAndUser(uint(id), userId)
	if err != nil {
		h.logger.Error("unable to get tag to database: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: err.Error()})
		return
	}

	// 3- Respond successful message
	h.logger.Debug("tag deleted successfully")
	utils.RespondWithJSON(w,
		&model.GenericResponse{Code: http.StatusOK, Status: true, Message: "Tag deleted successfully", Data: ""})
}
