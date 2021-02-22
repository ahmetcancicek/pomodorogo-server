package handler

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/dto"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/user"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// UserKey is used as a key for storing the User object in context at middleware
type UserKey struct{}

// UserUUIDKey is used as a key for storing the UserUUID in context at middleware
type UserUUIDKey struct{}

type AuthHandler struct {
	logger      *logrus.Logger
	UserService user.Service
	AuthService auth.Service
}

func NewAuthHandler(router *mux.Router, logger *logrus.Logger, userService user.Service, authService auth.Service) {
	handler := &AuthHandler{
		logger:      logger,
		UserService: userService,
		AuthService: authService,
	}
	router.HandleFunc("/api/v1/auth/signup", handler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/auth/signin", handler.SignIn).Methods(http.MethodPost)
	router.Use(handler.MiddlewareValidateUser)
}

func (h AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Refactoring

	// 1. Decode request body
	reqUser := r.Context().Value(UserKey{}).(model.User)

	// 2. Validate
	err := utils.PayloadValidator(reqUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.(validator.ValidationErrors).Error()}, w)
		return
	}

	// 3. Check if user exist in database
	_, err = h.UserService.FindByEmail(reqUser.Email)
	if err == nil {
		utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: model.ErrUserAlreadyExists}, w)
		return
	}

	// 4. Password Security
	hashedPass, err := h.hashPassword(reqUser.Password)
	tokenHash := utils.GenerateRandomString(15)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&model.GenericResponse{Status: false, Message: model.ErrUserSignUpFailed}, w)
		return
	}

	// 5. Create new user
	user := new(model.User)
	user.UUID = uuid.NewV4()
	user.FirstName = reqUser.FirstName
	user.LastName = reqUser.LastName
	user.Username = reqUser.Username
	user.Email = reqUser.Email
	user.Password = hashedPass
	user.TokenHash = tokenHash
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = h.UserService.Save(user)
	if err != nil {
		h.logger.Error("unable to insert user to database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrUserSignUpFailed}, w)
	}

	// 6- Respond successful message
	h.logger.Debug("user created successfully")
	w.WriteHeader(http.StatusCreated)
	utils.ToJSON(&model.GenericResponse{Code: 200, Status: true, Message: "User created successfully", Data: ""}, w)

}

func (h *AuthHandler) hashPassword(password string) (string, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//ah.logger.Error("unable to hash password", "error", err)
		return "", err
	}

	return string(hashedPass), nil
}

func (h AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Refactoring

	// 1. Decode request body
	reqUser := r.Context().Value(UserKey{}).(model.User)

	// 2. Validate
	err := utils.PayloadValidator(reqUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.(validator.ValidationErrors).Error()}, w)
		return
	}

	// 3. Authenticate
	user, err := h.UserService.FindByEmail(reqUser.Email)
	if err != nil {
		h.logger.Error("error fetching the user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrUserSignInFailed}, w)
		return
	}

	if valid := h.AuthService.Authenticate(reqUser.Password, user); !valid {
		h.logger.Debug("authentication of user failed")
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrUserSignInFailed}, w)
		return
	}

	// 4. Generate Access Token
	accessToken, err := h.AuthService.GenerateAccessToken(user)
	if err != nil {
		h.logger.Error("unable to generate access token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: "Unable to login the user. Please try again later"}, w)
		return
	}

	// 5. Generate Refresh Token
	refreshToken, err := h.AuthService.GenerateRefreshToken(user)
	if err != nil {
		h.logger.Error("unable to generate refresh token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: "Unable to login the user. Please try again later"}, w)
		return
	}

	// 6. Respond Token
	h.logger.Debug("successfully generated token: ", accessToken, refreshToken)
	w.WriteHeader(http.StatusOK)
	utils.ToJSON(&model.GenericResponse{Code: 200, Status: true, Message: "Successfully logged in", Data: &dto.AuthSignInResponseDTO{AccessToken: accessToken, RefreshToken: refreshToken}}, w)
}
