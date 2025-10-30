package register

type DTO struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type Response struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
