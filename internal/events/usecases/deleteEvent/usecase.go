package deleteEvent

import (
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/events/errors"
	userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type Request struct {
	Id string `json:"id" validate:"required,uuid"`
}

type Response struct {
	Event *domain.Event `json:"event"`
}

type DTO struct {
	Id   string
	User *userDomain.User
}

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	repository domain.Repository
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	event, err := usecase.repository.FindByID(dto.Id)

	if err != nil {
		return Response{}, errors.ErrEventNotFound
	}

	if event.CreatedByID != dto.User.ID {
		return Response{}, errors.ErrForbidden
	}

	if err := usecase.repository.Delete(event.ID.String()); err != nil {
		return Response{}, err
	}

	return Response{
		Event: event,
	}, nil
}

func NewUseCase(repository domain.Repository) UseCase {
	return &useCase{repository: repository}
}
