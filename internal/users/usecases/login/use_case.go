package login

import (
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	auth       *helper.Auth
	repository domain.Repository
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	if err := helper.Validate(dto); err != nil {
		return Response{}, err
	}

	user, err := usecase.repository.FindByEmail(dto.Email)

	if err != nil {
		return Response{}, domain.ErrUserNotFount
	}

	if !helper.VerifyPassword(dto.Password, user.Password) {
		return Response{}, domain.ErrInvalidPassword
	}

	token, err := usecase.auth.GenerateToken(user.ID, user.Email)

	if err != nil {
		return Response{}, domain.ErrGeneratingToken
	}

	return Response{
		ID:    user.ID.String(),
		Token: token,
	}, nil
}

func NewUseCase(repository domain.Repository, auth *helper.Auth) UseCase {
	return &useCase{repository: repository, auth: auth}
}
