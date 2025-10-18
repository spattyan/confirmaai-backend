package errors

import "errors"

var (
	ErrEventNotFound       = errors.New("event not found")
	ErrInvalidTimeFormat   = errors.New("invalid date and time format")
	ErrCreatingEventRole   = errors.New("error creating event role")
	ErrCreatingParticipant = errors.New("error creating participant")
	ErrForbidden           = errors.New("forbidden")
	ErrUnauthorizedAction  = errors.New("unauthorized action")
)
