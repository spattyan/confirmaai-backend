package createpermission

import userRepo "github.com/spattyan/confirmaai-backend/internal/users/domain"

type DTO struct {
	User *userRepo.User `json:"-"`
	Name string         `json:"name" validate:"required"`
	Key  string         `json:"key" validate:"required"`
}

type Response struct {
	ID string `json:"id"`
}
