package create

import (
	"errors"
	"time"

	"github.com/spattyan/confirmaai-backend/internal/events/domain"
)

type Request struct {
	Title            string `json:"title" validate:"required,min=2"`
	Description      string `json:"description"`
	Location         string `json:"location" validate:"required,min=4"`
	DateAndTime      string `json:"date_and_time" validate:"required"`
	ParticipantLimit int    `json:"participant_limit" validate:"required,min=1"`
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
}

type UseCase interface {
	Execute(dto DTO) (*domain.Event, error)
}

type useCase struct {
	repository domain.Repository
}

func (usecase *useCase) Execute(dto DTO) (*domain.Event, error) {

	parsedTime, err := time.Parse(time.DateTime, dto.DateAndTime)

	if err != nil {
		return nil, errors.New("invalid date and time format")
	}

	event := &domain.Event{
		Title:            dto.Title,
		Description:      dto.Description,
		Location:         dto.Location,
		DateAndTime:      parsedTime,
		ParticipantLimit: dto.ParticipantLimit,
	}

	if err := usecase.repository.Create(event); err != nil {
		return nil, err
	}

	return event, nil
}

func NewUseCase(repository domain.Repository) UseCase {
	return &useCase{repository: repository}
}
