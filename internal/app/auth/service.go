package auth

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
)

// AuthService represent the auth's service
type Service interface {
	Authenticate(password string, user *model.User) bool
	GenerateAccessToken(user *model.User) (string, error)
	GenerateRefreshToken(user *model.User) (string, error)
	GenerateCustomKey(userUUID string, password string) string
	ValidateAccessToken(token string) (string, error)
	ValidateRefreshToken(token string) (string, string, error)
}
