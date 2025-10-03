package login

import (
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/users/domain"
	"github.com/spattyan/confirmaai-backend/internal/users/errors"
)

type Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type Response struct {
	ID string `json:"id"`
}

type DTO struct {
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

	user, err := usecase.repository.FindByEmail(dto.Email)

	if err != nil {
		return Response{}, errors.ErrUserNotFount
	}

	if !helper.VerifyPassword(dto.Password, user.Password) {
		return Response{}, errors.ErrInvalidPassword
	}

	return Response{
		ID: user.ID.String(),
	}, nil
}

func NewUseCase(repository domain.Repository) UseCase {
	return &useCase{repository: repository}
}
