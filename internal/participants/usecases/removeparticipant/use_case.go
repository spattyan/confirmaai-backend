package removeparticipant

import (
	"github.com/spattyan/confirmaai-backend/helper"
	participantRepo "github.com/spattyan/confirmaai-backend/internal/participants/domain"
)

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	participantRepository participantRepo.Repository
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	if err := helper.Validate(dto); err != nil {
		return Response{}, err
	}

	if _, err := usecase.participantRepository.FindByID(dto.ParticipantID); err != nil {
		return Response{}, participantRepo.ErrParticipantNotFound
	}

	if err := usecase.participantRepository.Delete(dto.ParticipantID); err != nil {
		return Response{}, err
	}

	return Response{
		ID: dto.ParticipantID,
	}, nil
}

func NewUseCase(participantRepo participantRepo.Repository) UseCase {
	return &useCase{
		participantRepository: participantRepo,
	}
}
