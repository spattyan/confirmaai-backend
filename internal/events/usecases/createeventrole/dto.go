package createeventrole

import userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"

type DTO struct {
	EventID string           `json:"event_id" validate:"required,uuid"`
	Name    string           `json:"name" validate:"required,min=3,max=20"`
	Slots   int              `json:"slots" validate:"required,gte=1,lte=100"`
	User    *userDomain.User `json:"-"`
}

type Response struct {
	ID string `json:"id"`
}
