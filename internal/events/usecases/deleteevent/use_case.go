package deleteevent

import (
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
)

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	repository domain.Repository
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	if err := helper.Validate(dto); err != nil {
		return Response{}, err
	}

	event, err := usecase.repository.FindByID(dto.Id)

	if err != nil {
		return Response{}, domain.ErrEventNotFound
	}

	if event.CreatedByID != dto.User.ID {
		return Response{}, domain.ErrForbidden
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
