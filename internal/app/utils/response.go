package utils

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, response *model.GenericResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	_ = ToJSON(response, w)
}
