package handler

import (
	"encoding/json"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/dto"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/user"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/util"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthHandler struct {
	UserService user.Service
	AuthService auth.Service
}

func NewAuthHandler(router *mux.Router, userService user.Service, authService auth.Service) {
	handler := &AuthHandler{
		UserService: userService,
		AuthService: authService,
	}
	router.HandleFunc("/api/v1/auth/signup", handler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/auth/signin", handler.SignIn).Methods(http.MethodPost)
}

func (h AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Decode request body to dto object
	authSignUp := new(dto.AuthSignUpDTO)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&authSignUp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: "Invalid request payload"}, w)
		return
	}
	defer r.Body.Close()

	// 2. Validate
	err := util.PayloadValidator(authSignUp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.(validator.ValidationErrors).Error()}, w)
		return
	}

	// 3. Check if user exist in database
	_, err = h.UserService.FindByEmail(authSignUp.Email)
	if err == nil {
		util.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: model.ErrUserAlreadyExists}, w)
		return
	}

	// 4. Password Security
	hashedPass, err := h.hashPassword(authSignUp.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&model.GenericResponse{Status: false, Message: model.ErrUserSignUpFailed}, w)
		return
	}

	// 5. Create new user
	user := new(model.User)
	user.UUID = uuid.NewV4()
	user.FirstName = authSignUp.FirstName
	user.LastName = authSignUp.LastName
	user.Username = authSignUp.Username
	user.Email = authSignUp.Email
	user.Password = hashedPass
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = h.UserService.Save(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrUserSignUpFailed}, w)
	}

	// 6- Respond successful message
	w.WriteHeader(http.StatusCreated)
	util.ToJSON(&model.GenericResponse{Code: 200, Status: false, Message: "User created successfully", Data: ""}, w)

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

	// 1. Decode request body to dto object
	authSignIn := new(dto.AuthSignInDTO)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&authSignIn); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: "Invalid request payload"}, w)
	}
	defer r.Body.Close()

	// 2. Validate
	err := util.PayloadValidator(authSignIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.ToJSON(&model.GenericResponse{Code: http.StatusBadRequest, Status: false, Message: err.(validator.ValidationErrors).Error()}, w)
		return
	}

	// 3. Authenticate
	user, err := h.UserService.FindByEmail(authSignIn.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrUserSignInFailed}, w)
		return
	}

	if valid := h.AuthService.Authenticate(authSignIn.Password, user); !valid {
		w.WriteHeader(http.StatusBadRequest)
		util.ToJSON(&model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: model.ErrUserSignInFailed}, w)
		return
	}

	// 4. Generate Access Token
	accessToken, err := h.AuthService.GenerateAccessToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: "Unable to login the user. Please try again later"}, w)
	}

	// 5. Generate Refresh Token
	refreshToken, err := h.AuthService.GenerateRefreshToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(model.GenericResponse{Code: http.StatusInternalServerError, Status: false, Message: "Unable to login the user. Please try again later"}, w)
	}

	// 6. Respond Token
	w.WriteHeader(http.StatusOK)
	util.ToJSON(&model.GenericResponse{Code: 200, Status: false, Message: "Successfully logged in", Data: &dto.AuthSignInResponseDTO{AccessToken: accessToken, RefreshToken: refreshToken}}, w)
}
