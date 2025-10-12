package getById

import (
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
)

type Request struct {
	Id string `json:"id" validate:"required,uuid"`
}

type Response struct {
	Event EventResponse `json:"event"`
}

type EventResponse struct {
	Title            string               `json:"title"`
	Description      string               `json:"description"`
	Location         string               `json:"location"`
	DateAndTime      string               `json:"date_and_time"`
	ParticipantLimit int                  `json:"participant_limit"`
	CreatedBy        string               `json:"created_by"`
	Participants     ParticipantsResponse `json:"participants"`
}

type ParticipantsResponse struct {
	Count int                        `json:"count"`
	List  []ParticipantsListResponse `json:"list"`
}

type ParticipantsListResponse struct {
	ID string `json:"id"`
}

type DTO struct {
	Id string
}

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	repository domain.Repository
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	event, err := usecase.repository.FindByID(dto.Id)

	if err != nil {
		return Response{}, err
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
