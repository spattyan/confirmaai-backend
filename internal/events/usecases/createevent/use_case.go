package createevent

import (
	"fmt"
	"time"

	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/createeventrole"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/deleteevent"
	"github.com/spattyan/confirmaai-backend/internal/participants/usecases/createparticipant"
)

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	repository        domain.Repository
	createEventRole   createeventrole.UseCase
	createParticipant createparticipant.UseCase
	deleteEvent       deleteevent.UseCase
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	if err := helper.Validate(dto); err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	parsedTime, err := time.Parse(time.RFC3339, dto.DateAndTime)

	if err != nil {
		return Response{}, domain.ErrInvalidTimeFormat
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

	eventRole, err := usecase.createEventRole.Execute(createeventrole.DTO{
		EventID: event.ID.String(),
		Name:    "Organizer",
		Slots:   1,
		User:    dto.User,
	})

	if err != nil {
		return Response{}, domain.ErrCreatingEventRole
	}

	_, err = usecase.createParticipant.Execute(createparticipant.DTO{
		EventID: event.ID.String(),
		UserID:  dto.User.ID.String(),
		RoleID:  eventRole.ID,
	})

	if err != nil {
		// TODO: call delete event role usecase

		_, err := usecase.deleteEvent.Execute(deleteevent.DTO{
			Id:   event.ID.String(),
			User: dto.User,
		})
		if err != nil {
			return Response{}, err
		}

		return Response{}, domain.ErrCreatingParticipant
	}

	return Response{
		ID: event.ID.String(),
	}, nil
}

func NewUseCase(repository domain.Repository, createEventRole createeventrole.UseCase, createParticipant createparticipant.UseCase) UseCase {
	return &useCase{repository: repository, createEventRole: createEventRole, createParticipant: createParticipant}
}
