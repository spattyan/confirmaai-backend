package list

import (
	"time"

	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	domain2 "github.com/spattyan/confirmaai-backend/internal/participants/domain"
)

type Request struct {
}

type Response struct {
	Count  int              `json:"count"`
	Events []ResponseObject `json:"events"`
}
type ResponseObject struct {
	Title            string                `json:"title"`
	Description      string                `json:"description"`
	Location         string                `json:"location"`
	DateAndTime      string                `json:"date_and_time"`
	ParticipantLimit int                   `json:"participant_limit"`
	Participants     []domain2.Participant `json:"participants"`
}

type DTO struct {
}

type UseCase interface {
	Execute() (Response, error)
}

type useCase struct {
	repository domain.Repository
}

func (usecase *useCase) Execute() (Response, error) {

	events, err := usecase.repository.List()

	if err != nil {
		return Response{}, err
	}

	responseObjects := make([]ResponseObject, len(events))
	for i, event := range events {

		participants, err := usecase.repository.ListParticipantsByEventID(event.ID.String())
		if err != nil {
			return Response{}, err
		}

		event.Participants = participants
		responseObjects[i] = ResponseObject{
			Title:            event.Title,
			Description:      event.Description,
			Location:         event.Location,
			DateAndTime:      event.DateAndTime.Format(time.RFC3339),
			ParticipantLimit: event.ParticipantLimit,
			Participants:     event.Participants,
		}
	}

	responseBody := Response{
		Count:  len(events),
		Events: responseObjects,
	}

	return responseBody, nil
}

func NewUseCase(repository domain.Repository) UseCase {
	return &useCase{repository: repository}
}
