package createeventrole

import (
	"github.com/google/uuid"
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

	event, err := usecase.repository.FindByID(dto.EventID)

	if err != nil {
		return Response{}, domain.ErrEventNotFound
	}

	if event.CreatedByID != dto.User.ID {
		return Response{}, domain.ErrUnauthorizedAction
	}

	role := &domain.EventRole{
		EventID: uuid.MustParse(dto.EventID),
		Name:    dto.Name,
		Slots:   dto.Slots,
	}

	if err := usecase.repository.CreateEventRole(role); err != nil {
		return Response{}, err
	}

	return Response{
		ID: role.ID.String(),
	}, nil
}

func NewUseCase(repository domain.Repository) UseCase {
	return &useCase{repository: repository}
}
