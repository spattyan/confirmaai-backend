package register

import (
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	repository domain.Repository
	auth       *helper.Auth
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	if err := helper.Validate(dto); err != nil {
		return Response{}, err
	}

	hash, err := helper.HashPassword(dto.Password)

	if err != nil {
		return Response{}, domain.ErrHashingPassword
	}

	user := &domain.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hash,
	}

	if err, _ := usecase.repository.FindByEmail(user.Email); err != nil {
		return Response{}, domain.ErrUserAlreadyExists
	}

	if err := usecase.repository.Create(user); err != nil {
		return Response{}, err
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
