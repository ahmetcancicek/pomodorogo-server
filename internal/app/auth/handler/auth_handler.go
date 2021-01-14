package handler

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/gorilla/mux"
	"net/http"
)

type AuthHandler struct {
	UserService auth.Service
}

func NewAuthHandler(router *mux.Router, userService auth.Service) {
	handler := &AuthHandler{
		UserService: userService,
	}
	router.HandleFunc("/api/v1/auth/signup", handler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/auth/signin", handler.SignIn).Methods(http.MethodPost)
}

func (h AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {

}
