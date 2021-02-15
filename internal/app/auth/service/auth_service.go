package service

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authService struct {
}

// NewAuthService will create new an useService object representation of of auth.Service interface
func NewAuthService() auth.Service {
	return &authService{}
}

func (a authService) Authenticate(password string, user *model.User) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		//auth.logger.Debug("password hashes are not same")
		return false
	}
	return true
}

// AccessTokenCustomClaims specifies the claims for access token
type AccessTokenCustomClaims struct {
	UserID  string
	KeyType string
	jwt.StandardClaims
}

func (a authService) GenerateAccessToken(user *model.User) (string, error) {
	// TODO: We must get secret key from config file
	accessSecret := []byte("AllYourBase")
	accessTokenExpireDuration := time.Duration(time.Minute * 5)

	accessTokenExpiresTime := time.Now().Add(accessTokenExpireDuration)
	accessTokenUUID := uuid.NewV4()

	// Create Access Token
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["user_uuid"] = user.UUID.String()
	accessTokenClaims["uuid"] = accessTokenUUID
	accessTokenClaims["exp"] = accessTokenExpiresTime.Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	return accessToken.SignedString([]byte(accessSecret))

}

func (a authService) GenerateRefreshToken(user *model.User) (string, error) {
	// TODO: We must get secret key from config file
	accessSecret := []byte("AllYourBase")
	refreshTokenExpireDuration := time.Duration(time.Minute * 5)

	refreshTokenExpiresTime := time.Now().Add(refreshTokenExpireDuration)
	refreshTokenUUID := uuid.NewV4()

	// Create Access Token
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["user_uuid"] = user.UUID.String()
	refreshTokenClaims["exp"] = refreshTokenExpiresTime.Unix()
	refreshTokenClaims["uuid"] = refreshTokenUUID
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	return refreshToken.SignedString([]byte(accessSecret))
}

func (a authService) ValidateAccessToken(token string) (string, error) {
	panic("implement me")
}

func (a authService) ValidateRefreshToken(token string) (string, error) {
	panic("implement me")
}
