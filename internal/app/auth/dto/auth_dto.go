package dto

type AuthSignUpDTO struct {
	FirstName string `validate:"required" json:"firstName"`
	LastName  string `validate:"required" json:"lastName"`
	Username  string `validate:"required" json:"username"`
	Email     string `validate:"required" json:"email"`
	Password  string `validate:"required" json:"password"`
}

type AuthSignInDTO struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type AuthSignInResponseDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
