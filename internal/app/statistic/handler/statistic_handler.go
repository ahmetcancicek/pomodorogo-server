package handler

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/handler"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/statistic"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/statistic/dto"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type StatisticHandler struct {
	logger           *logrus.Logger
	StatisticService statistic.Service
}

func NewStatisticHandler(r *mux.Router, log *logrus.Logger, statService statistic.Service, mf mux.MiddlewareFunc) {
	statHandler := &StatisticHandler{
		logger:           log,
		StatisticService: statService,
	}

	r.HandleFunc("/api/v1/statistics", statHandler.create).Methods(http.MethodPost)
	r.Use(mf)
}

func (h StatisticHandler) create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statDTO := new(dto.StatisticDTO)

	// 1. Decode request body
	err := utils.FromJSON(statDTO, r.Body)
	if err != nil {
		h.logger.Error("unable to decode work duration json", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.Error()}, w)
		return
	}
	defer r.Body.Close()

	// 2. Save
	userId := r.Context().Value(handler.UserIDKey{}).(uint)
	statDTO, err = h.StatisticService.Save(statDTO, userId)
	if err != nil {
		h.logger.Error("unable to insert work duration to database: ", err)
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: err.Error()}, w)
		return
	}

	// 3. Respond successful message
	h.logger.Debug("work duration added successfully")
	w.WriteHeader(http.StatusCreated)
	utils.ToJSON(&model.GenericResponse{Code: 200, Status: true, Message: "Work duration added successfully", Data: statDTO}, w)

}
