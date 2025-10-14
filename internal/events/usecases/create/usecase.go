package create

import (
	"time"

	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/events/errors"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/createEventRole"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/deleteEvent"
	"github.com/spattyan/confirmaai-backend/internal/participants/usecases/createParticipant"
	userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type Request struct {
	Title            string `json:"title" validate:"required,min=2"`
	Description      string `json:"description"`
	Location         string `json:"location" validate:"required,min=4"`
	DateAndTime      string `json:"date_and_time" validate:"required"`
	ParticipantLimit int    `json:"participant_limit" validate:"required,gte=0"`
}

type Response struct {
	ID string `json:"id"`
}

type DTO struct {
	Title            string
	Description      string
	Location         string
	DateAndTime      string
	ParticipantLimit int
	User             *userDomain.User
}

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	repository        domain.Repository
	createEventRole   createEventRole.UseCase
	createParticipant createParticipant.UseCase
	deleteEvent       deleteEvent.UseCase
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	_, err := helper.Validate(Request{
		Title:            dto.Title,
		Description:      dto.Description,
		Location:         dto.Location,
		DateAndTime:      dto.DateAndTime,
		ParticipantLimit: dto.ParticipantLimit,
	})

	if err != nil {
		return Response{}, err
	}

	parsedTime, err := time.Parse(time.DateTime, dto.DateAndTime)

	if err != nil {
		return Response{}, errors.ErrInvalidTimeFormat
	}

	event := &domain.Event{
		Title:            dto.Title,
		Description:      dto.Description,
		Location:         dto.Location,
		DateAndTime:      parsedTime,
		ParticipantLimit: dto.ParticipantLimit,
		CreatedByID:      dto.User.ID,
	}

	if err := usecase.repository.Create(event); err != nil {
		return Response{}, err
	}

	eventRole, err := usecase.createEventRole.Execute(createEventRole.DTO{
		EventID: event.ID.String(),
		Name:    "Organizer",
		Slots:   1,
	})

	if err != nil {
		return Response{}, errors.ErrCreatingEventRole
	}

	_, err = usecase.createParticipant.Execute(createParticipant.DTO{
		EventID: event.ID.String(),
		UserID:  dto.User.ID.String(),
		RoleID:  eventRole.ID,
	})

	if err != nil {
		// TODO: call delete event role usecase

		_, err := usecase.deleteEvent.Execute(deleteEvent.DTO{
			Id:   event.ID.String(),
			User: dto.User,
		})
		if err != nil {
			return Response{}, err
		}

		return Response{}, errors.ErrCreatingParticipant
	}

	return Response{
		ID: event.ID.String(),
	}, nil
}

func NewUseCase(repository domain.Repository, createEventRole createEventRole.UseCase, createParticipant createParticipant.UseCase) UseCase {
	return &useCase{repository: repository, createEventRole: createEventRole, createParticipant: createParticipant}
}
