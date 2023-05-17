package constants

import "errors"

var (
	ErrUserNotFound       = errors.New("User not found")
	ErrPasswordNotMatch   = errors.New("Password is not match")
	ErrRecordNotFound     = errors.New("record not found")
	ErrNamaTokoIsRequired = errors.New("Nama toko is required")
)
