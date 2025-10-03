package errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrHashingPassword   = errors.New("error hashing password")
	ErrUserNotFount      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
)
