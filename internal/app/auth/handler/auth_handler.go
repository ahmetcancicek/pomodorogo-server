package handler

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/account"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/dto"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// UserKey is used as a key for storing the User object in context at middleware
type UserKey struct{}

// UserUUIDKey is used as a key for storing the UserUUID in context at middleware
type UserUUIDKey struct{}

// UserIDKey is used as a key for storing the UserIDKey in context at middleware
type UserIDKey struct{}

type AuthHandler struct {
	logger         *logrus.Logger
	AccountService account.Service
	AuthService    auth.Service
}

func NewAuthHandler(r *mux.Router, log *logrus.Logger, us account.Service, as auth.Service) *AuthHandler {
	authHandler := &AuthHandler{
		logger:         log,
		AccountService: us,
		AuthService:    as,
	}
	r.HandleFunc("/api/v1/auth/signup", authHandler.signUp).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/auth/signin", authHandler.signIn).Methods(http.MethodPost)
	r.Use(authHandler.MiddlewareValidateUser)
	return authHandler
}

func (h AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {

	// 1. Decode request body
	user := r.Context().Value(UserKey{}).(model.User)

	// 2. Password Security
	hashedPass, err := h.hashPassword(user.Password)
	tokenHash := utils.GenerateRandomString(15)
	if err != nil {
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrUserSignUpFailed})
		return
	}

	// 3. Save
	user.Password = hashedPass
	user.TokenHash = tokenHash
	err = h.AccountService.Save(&user)
	if err != nil {
		h.logger.Error("unable to insert user to database: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: err.Error()})
		return
	}

	// 4- Respond successful message
	h.logger.Debug("user created successfully")
	utils.RespondWithJSON(w,
		&model.GenericResponse{Code: http.StatusCreated, Status: true, Message: "User created successfully", Data: ""})
}

func (h *AuthHandler) hashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		h.logger.Error("unable to hash password", "error", err)
		return "", err
	}
	return string(hashedPass), nil
}

func (h AuthHandler) signIn(w http.ResponseWriter, r *http.Request) {

	// 1. Decode request body
	reqUser := r.Context().Value(UserKey{}).(model.User)

	// 2. Validate
	err := utils.PayloadValidator(reqUser)
	if err != nil {
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.(validator.ValidationErrors).Error()})
		return
	}

	// 3. Authenticate
	user, err := h.AccountService.FindByEmail(reqUser.Email)
	if err != nil {
		h.logger.Error("error fetching the user: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrUserSignInFailed})
		return
	}

	if valid := h.AuthService.Authenticate(reqUser.Password, user); !valid {
		h.logger.Debug("authentication of user failed")
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrUserSignInFailed})
		return
	}

	// 4. Generate Access Token
	accessToken, err := h.AuthService.GenerateAccessToken(user)
	if err != nil {
		h.logger.Error("unable to generate access token: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: "Unable to login the user. Please try again later"})

		return
	}

	// 5. Generate Refresh Token
	refreshToken, err := h.AuthService.GenerateRefreshToken(user)
	if err != nil {
		h.logger.Error("unable to generate refresh token: ", err)
		utils.RespondWithJSON(w,
			&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: "Unable to login the user. Please try again later"})
		return
	}

	// 6. Respond Token
	h.logger.Debug("successfully generated token: ", accessToken, refreshToken)
	utils.RespondWithJSON(w,
		&model.GenericResponse{
			Code:    http.StatusOK,
			Status:  true,
			Message: "Successfully logged in",
			Data:    &dto.AuthSignInResponseDTO{AccessToken: accessToken, RefreshToken: refreshToken},
		})
}
