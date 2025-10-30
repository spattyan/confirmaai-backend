package haspermission

import (
	"github.com/spattyan/confirmaai-backend/helper"
	participantRepo "github.com/spattyan/confirmaai-backend/internal/participants/domain"
)

type UseCase interface {
	Execute(dto DTO) bool
}

type useCase struct {
	participantRepository participantRepo.Repository
}

func (useCase *useCase) Execute(dto DTO) bool {

	if err := helper.Validate(dto); err != nil {
		return false
	}

	permission, err := useCase.participantRepository.GetPermissionByKey(dto.Permission)

	if err != nil {
		return false
	}

	participant, err := useCase.participantRepository.FindByID(dto.ParticipantID)

	if err != nil {
		return false
	}

	for _, p := range participant.Permissions {
		if p.ID == permission.ID {
			return true
		}
	}

	return false

}

func NewUseCase(participantRepo participantRepo.Repository) UseCase {
	return &useCase{

		participantRepository: participantRepo,
	}
}
