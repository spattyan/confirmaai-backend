package deleteevent

import (
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type DTO struct {
	Id   string           `json:"id" validate:"required,uuid"`
	User *userDomain.User `json:"-"`
}

type Response struct {
	Event *domain.Event `json:"event"` // need to change this to appropriate response fields.
}
