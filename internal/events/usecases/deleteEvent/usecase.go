package deleteEvent

import (
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/events/errors"
)

type Request struct {
	Id string `json:"id" validate:"required,uuid"`
}

type Response struct {
	Event *domain.Event `json:"event"`
}

type DTO struct {
	Id string
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
