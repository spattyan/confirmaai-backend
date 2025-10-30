package listevents

import (
	"time"

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

	events, err := usecase.repository.List()

	if err != nil {
		return Response{}, err
	}

	responseObjects := make([]ResponseObject, len(events))
	for i, event := range events {

		responseObjects[i] = ResponseObject{
			Title:            event.Title,
			Description:      event.Description,
			Location:         event.Location,
			DateAndTime:      event.DateAndTime.Format(time.RFC3339),
			ParticipantLimit: event.ParticipantLimit,
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
