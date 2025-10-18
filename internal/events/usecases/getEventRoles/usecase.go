package getEventRoles

import (
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
)

type Request struct {
	Id string `json:"id" validate:"required,uuid"`
}

type Response struct {
	Roles []RolesResponse `json:"roles"`
}

type RolesResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Slots int    `json:"slots"`
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

	roles, err := usecase.repository.ListEventRolesByEventID(dto.Id)

	if err != nil {
		return Response{}, err
	}
	rolesResponse := make([]RolesResponse, len(roles))
	for i, role := range roles {
		rolesResponse[i] = RolesResponse{
			ID:    role.ID.String(),
			Name:  role.Name,
			Slots: role.Slots,
		}
	}

	return Response{
		Roles: rolesResponse,
	}, nil
}

func NewUseCase(repository domain.Repository) UseCase {
	return &useCase{repository: repository}
}
