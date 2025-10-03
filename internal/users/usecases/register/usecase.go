package register

import (
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/users/domain"
	"github.com/spattyan/confirmaai-backend/internal/users/errors"
)

type Request struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type Response struct {
	ID string `json:"id"`
}

type DTO struct {
	Name     string
	Email    string
	Password string
}

type UseCase interface {
	Execute(dto DTO) (Response, error)
}

type useCase struct {
	repository domain.Repository
}

func (usecase *useCase) Execute(dto DTO) (Response, error) {

	hash, err := helper.HashPassword(dto.Password)

	if err != nil {
		return Response{}, errors.ErrHashingPassword
	}

	user := &domain.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hash,
	}

	if err, _ := usecase.repository.FindByEmail(user.Email); err != nil {
		return Response{}, errors.ErrUserAlreadyExists
	}

	if err := usecase.repository.Create(user); err != nil {
		return Response{}, err
	}

	return Response{
		ID: user.ID.String(),
	}, nil
}

func NewUseCase(repository domain.Repository) UseCase {
	return &useCase{repository: repository}
}
