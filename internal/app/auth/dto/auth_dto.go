package dto

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type AuthSignUpDTO struct {
	FirstName string `validate:"required" json:"firstName"`
	LastName  string `validate:"required" json:"lastName"`
	Username  string `validate:"required" json:"username"`
	Email     string `validate:"required" json:"email"`
	Password  string `validate:"required" json:"password"`
}

type AuthLoginDTO struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type AuthLoginResponseDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenDTO struct {
	AccessToken             string `json:"accessToken"`
	RefreshToken            string `json:"refreshToken"`
	AccessTokenExpiresTime  time.Time
	RefreshTokenExpiresTime time.Time
	AccessTokenUUID         uuid.UUID
	RefreshTokenUUID        uuid.UUID
}
