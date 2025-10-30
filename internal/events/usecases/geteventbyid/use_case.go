package geteventbyid

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

	participantsList := make([]ParticipantsListResponse, len(event.Participants))
	for i, participant := range event.Participants {
		participantsList[i] = ParticipantsListResponse{
			ID: participant.ID.String(),
		}
	}
	event.Participants = nil

	eventResponse := EventResponse{
		Title:            event.Title,
		Description:      event.Description,
		Location:         event.Location,
		DateAndTime:      event.DateAndTime.Format("2006-01-02 15:04:05"),
		ParticipantLimit: event.ParticipantLimit,
		CreatedBy:        event.CreatedByID.String(),
		Participants: ParticipantsResponse{
			Count: len(participantsList),
			List:  participantsList,
		},
	}

	return Response{
		Event: eventResponse,
	}, nil
}

func NewUseCase(repository domain.Repository) UseCase {
	return &useCase{repository: repository}
}
