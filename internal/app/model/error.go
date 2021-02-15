package model

import (
	"fmt"
)

var (
	ErrUserSignInFailed  = fmt.Sprint("No user account exists with given email or password.")
	ErrUserSignUpFailed  = fmt.Sprintf("Unable to create user. Please try again later")
	ErrUserNotFound      = fmt.Sprintf("No user account exists with given email. Please sign in first")
	ErrUserAlreadyExists = fmt.Sprintf("User already exists with the given email")
)
