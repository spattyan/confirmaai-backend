package createpermission

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/helper"
	participantRepo "github.com/spattyan/confirmaai-backend/internal/participants/domain"
	userRepo "github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	userRepository        userRepo.Repository
	participantRepository participantRepo.Repository
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	if err := helper.Validate(dto); err != nil {
		return Response{}, err
	}

	if !dto.User.Admin {
		return Response{}, participantRepo.ErrUserNotAdmin
	}

	err := usecase.participantRepository.CreatePermission(dto.Name, dto.Key)

	if err != nil {
		return Response{}, fmt.Errorf("failed to createevent permission: %w", err)
	}

	return Response{
		ID: uuid.New().String(),
	}, nil
}

func NewUseCase(userRepository userRepo.Repository, participantRepo participantRepo.Repository) UseCase {
	return &useCase{
		userRepository:        userRepository,
		participantRepository: participantRepo,
	}
}
