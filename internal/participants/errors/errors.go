package errors

import "errors"

var (
	ErrParticipantNotFound      = errors.New("participant not found")
	ErrEventNotFound            = errors.New("event not found")
	ErrUserNotFound             = errors.New("user not exists")
	ErrParticipantAlreadyExists = errors.New("participant already exists")
)
