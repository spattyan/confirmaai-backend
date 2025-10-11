package createParticipant

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/helper"
	eventRepo "github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/participants/domain"
	participantRepo "github.com/spattyan/confirmaai-backend/internal/participants/domain"
	"github.com/spattyan/confirmaai-backend/internal/participants/errors"
	userRepo "github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type Request struct {
	EventID string `json:"event_id" validate:"required,uuid"`
	UserID  string `json:"user_id" validate:"required,uuid"`
	RoleID  string `json:"role_id" validate:"required,uuid"`
}

type Response struct {
	ID string `json:"id"`
}

type DTO struct {
	EventID string
	UserID  string
	RoleID  string
}

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	userRepository        userRepo.Repository
	eventRepository       eventRepo.Repository
	participantRepository participantRepo.Repository
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	_, err := helper.Validate(Request{
		EventID: dto.EventID,
		UserID:  dto.UserID,
		RoleID:  dto.RoleID,
	})

	if err != nil {
		return Response{}, err
	}

	if _, err := usecase.eventRepository.FindByID(dto.EventID); err != nil {
		return Response{}, errors.ErrEventNotFound
	}

	if _, err := usecase.userRepository.FindByID(dto.UserID); err != nil {
		fmt.Println(err)
		return Response{}, errors.ErrUserNotFound
	}

	if _, err := usecase.participantRepository.FindByEventAndUser(dto.EventID, dto.UserID); err == nil {
		return Response{}, errors.ErrParticipantAlreadyExists
	}

	participant := &domain.Participant{
		EventID: uuid.MustParse(dto.EventID),
		UserID:  uuid.MustParse(dto.UserID),
		RoleID:  uuid.MustParse(dto.RoleID),
	}

	if err := usecase.participantRepository.Create(participant); err != nil {
		return Response{}, err
	}

	return Response{
		ID: participant.ID.String(),
	}, nil
}

func NewUseCase(userRepository userRepo.Repository, eventsRepository eventRepo.Repository, participantRepo participantRepo.Repository) UseCase {
	return &useCase{
		userRepository:        userRepository,
		eventRepository:       eventsRepository,
		participantRepository: participantRepo,
	}
}
