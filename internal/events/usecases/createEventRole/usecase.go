package createEventRole

import (
	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/events/errors"
)

type Request struct {
	EventID string `json:"event_id" validate:"required,uuid"`
	Name    string `json:"name" validate:"required"`
	Slots   int    `json:"slots" validate:"required,min=1,max=512"`
}

type Response struct {
	ID string `json:"id"`
}

type DTO struct {
	EventID string
	Name    string
	Slots   int
}

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	repository domain.Repository
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	if _, err := helper.Validate(Request{
		EventID: dto.EventID,
		Name:    dto.Name,
		Slots:   dto.Slots,
	}); err != nil {
		return Response{}, err
	}

	if _, err := usecase.repository.FindByID(dto.EventID); err != nil {
		return Response{}, errors.ErrEventNotFound
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
