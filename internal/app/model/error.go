package model

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("your requested item is not found")

	ErrUserSignInFailed  = fmt.Sprint("No user account exists with given email or password.")
	ErrUserSignUpFailed  = fmt.Sprintf("Unable to create user. Please try again later")
	ErrUserAlreadyExists = fmt.Sprintf("User already exists with the given email")

	ErrTagCreateFailed = fmt.Sprintf("Unable to tag user. Please try again later")
)
