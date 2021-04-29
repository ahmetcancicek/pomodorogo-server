package handler

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/handler"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/pomodoro"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/pomodoro/dto"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PomodoroHandler struct {
	logger          *logrus.Logger
	PomodoroService pomodoro.Service
}

func NewStatisticHandler(r *mux.Router, log *logrus.Logger, pomodoroService pomodoro.Service, mf mux.MiddlewareFunc) {
	pomodoroHandler := &PomodoroHandler{
		logger:          log,
		PomodoroService: pomodoroService,
	}

	r.HandleFunc("/api/v1/pomodoro", pomodoroHandler.create).Methods(http.MethodPost)
	r.Use(mf)
}

func (h PomodoroHandler) create(w http.ResponseWriter, r *http.Request) {

	pomodoroDTO := new(dto.PomodoroDTO)

	// 1. Decode request body
	err := utils.FromJSON(pomodoroDTO, r.Body)
	if err != nil {
		h.logger.Error("unable to decode work duration json", err.Error())
		utils.RespondWithJSON(
			w,
			&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.Error()},
		)
		return
	}
	defer r.Body.Close()

	// 2. Save
	userId := r.Context().Value(handler.UserIDKey{}).(uint)
	createdPomodoro, err := h.PomodoroService.Save(dto.ToPomodoro(pomodoroDTO), userId)
	if err != nil {
		h.logger.Error("unable to insert work duration to database: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: err.Error()})
		return
	}

	// 3. Respond successful message
	h.logger.Debug("work duration added successfully")
	utils.RespondWithJSON(w,
		&model.GenericResponse{Code: http.StatusCreated, Status: true, Message: "Work duration added successfully", Data: dto.ToPomodoroDTO(createdPomodoro)})
}
