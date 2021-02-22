package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/auth"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"time"
)

type authService struct {
	configs *utils.Configurations
}

// NewAuthService will create new an useService object representation of of auth.Service interface
func NewAuthService(c *utils.Configurations) auth.Service {
	return &authService{configs: c}
}

func (a authService) Authenticate(password string, user *model.User) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		//auth.logger.Debug("password hashes are not same")
		return false
	}
	return true
}

type RefreshTokenCustomClaims struct {
	UserUUID  string
	CustomKey string
	KeyType   string
	jwt.StandardClaims
}

type AccessTokenCustomClaims struct {
	UserUUID string
	KeyType  string
	jwt.StandardClaims
}

// GenerateAccessToken generates a new access token for the given user
func (a authService) GenerateAccessToken(user *model.User) (string, error) {
	tokenType := "access"

	claims := AccessTokenCustomClaims{
		user.UUID.String(),
		tokenType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(a.configs.JwtExpiration)).Unix(),
			Issuer:    "pomodorogo.auth.service",
		},
	}

	signBytes, err := ioutil.ReadFile(a.configs.AccessTokenPrivateKeyPath)
	if err != nil {
		fmt.Print("Error")
		return "", errors.New("could not generate access token. please try again later")
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", errors.New("could not generate access token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)

}

// GenerateRefreshToken generate a new refresh token for the given user
func (a authService) GenerateRefreshToken(user *model.User) (string, error) {

	customKey := a.GenerateCustomKey(user.UUID.String(), user.TokenHash)
	tokenType := "refresh"

	claims := RefreshTokenCustomClaims{
		user.UUID.String(),
		customKey,
		tokenType,
		jwt.StandardClaims{
			Issuer: "pomodorogo.auth.service",
		},
	}

	signBytes, err := ioutil.ReadFile(a.configs.RefreshTokenPrivateKeyPath)
	if err != nil {
		return "", errors.New("could not generate refresh token. please try again later")
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", errors.New("could not generate refresh token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

// GenerateCustomKey creates a new key for our jwt payload
func (a *authService) GenerateCustomKey(userUUID string, tokenHash string) string {
	h := hmac.New(sha256.New, []byte(tokenHash))
	h.Write([]byte(userUUID))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

//
func (a authService) ValidateAccessToken(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected signing method in auth token")
		}

		verifyBytes, err := ioutil.ReadFile(a.configs.AccessTokenPublicKeyPath)
		if err != nil {
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*AccessTokenCustomClaims)
	if !ok || !token.Valid || claims.UserUUID == "" || claims.KeyType != "access" {
		return "", errors.New("invalid token: authentication failed")
	}

	return claims.UserUUID, nil
}

//
func (a authService) ValidateRefreshToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected signing method in auth token")
		}

		verifyBytes, err := ioutil.ReadFile(a.configs.RefreshTokenPublicKeyPath)
		if err != nil {
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*RefreshTokenCustomClaims)
	if !ok || !token.Valid || claims.UserUUID == "" || claims.KeyType != "refresh" {
		return "", "", errors.New("invalid token: authentication failed")
	}
	return claims.UserUUID, claims.CustomKey, nil
}
