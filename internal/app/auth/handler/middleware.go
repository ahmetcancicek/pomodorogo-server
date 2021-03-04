package handler

import (
	"context"
	"errors"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// MiddlewareValidateUser validates the user in the request
func (h *AuthHandler) MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		user := &model.User{}

		err := utils.FromJSON(user, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&model.GenericResponse{Status: false, Message: err.Error()}, w)
			return
		}

		// 2. Validate
		err = utils.PayloadValidator(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.(validator.ValidationErrors).Error()}, w)
			return
		}

		// Add the user to the context
		ctx := context.WithValue(r.Context(), UserKey{}, *user)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

//
func (h *AuthHandler) MiddlewareValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		token, err := extractToken(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: "Authentication failed. Token not provided"}, w)
			return
		}

		userUUID, err := h.AuthService.ValidateAccessToken(token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: "Authentication failed. Invalid token"}, w)
			return
		}

		user, err := h.AccountService.FindByUUID(userUUID)

		ctx := context.WithValue(r.Context(), UserIDKey{}, user.ID)
		//ctx := context.WithValue(r.Context(), UserUUIDKey{}, userUUID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

//
func (h *AuthHandler) MiddlewareValidateRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		token, err := extractToken(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: "Authentication failed. Token not provided"}, w)
			return
		}

		userUUID, customKey, err := h.AuthService.ValidateRefreshToken(token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: "Authentication failed. Invalid token"}, w)
			return
		}

		user, err := h.AccountService.FindByUUID(userUUID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: "Authentication failed. Invalid token"}, w)
			return
		}

		actualCustomKey := h.AuthService.GenerateCustomKey(user.UUID.String(), user.TokenHash)
		if customKey != actualCustomKey {
			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: "Authentication failed. Invalid token"}, w)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey{}, *user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		return "", errors.New("token not provided")
	}
	return authHeaderContent[1], nil
}
