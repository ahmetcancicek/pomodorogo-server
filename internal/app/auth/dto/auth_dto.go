package dto

type AuthSignUpDTO struct {
	FirstName string `validate:"required,max=50,min=3" json:"firstName"`
	LastName  string `validate:"required,max=50,min=3" json:"lastName"`
	Username  string `validate:"required,max=50,min=3" json:"username"`
	Email     string `validate:"required,max=75,min=3" json:"email"`
	Password  string `validate:"required,max=50,min=3" json:"password"`
}

type AuthSignInDTO struct {
	Email    string `validate:"required,max=75,min=3" json:"email"`
	Password string `validate:"required,max=50,min=3" json:"password"`
}

type AuthSignInResponseDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
