package constants

import "errors"

var (
	ErrUserNotFound     = errors.New("User not found")
	ErrPasswordNotMatch = errors.New("Password is not match")
)
