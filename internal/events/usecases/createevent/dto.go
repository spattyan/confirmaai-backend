package createevent

import userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"

type Response struct {
	ID string `json:"id"`
}

type DTO struct {
	Title            string `json:"title" validate:"required,min=2,max=60"`
	Description      string `json:"description" validate:"min=2,max=2048"`
	Location         string `json:"location" validate:"required,min=4,max=100"`
	DateAndTime      string `json:"date_and_time" validate:"required,datetime"`
	ParticipantLimit int    `json:"participant_limit" validate:"required,gte=0,lte=512"`
	User             *userDomain.User
}
