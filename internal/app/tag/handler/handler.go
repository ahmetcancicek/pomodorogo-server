package handler

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/tag"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type TagHandler struct {
	logger     *logrus.Logger
	TagService tag.Service
}

func NewTagHandler(router *mux.Router, logger *logrus.Logger, tagService tag.Service, middlewareFunc mux.MiddlewareFunc) {
	tagHandler := &TagHandler{
		logger:     logger,
		TagService: tagService,
	}

	router.HandleFunc("/api/v1/tag/create", tagHandler.Create).Methods(http.MethodPost)
	router.Use(middlewareFunc)
}

func (h TagHandler) Create(w http.ResponseWriter, r *http.Request) {

}
