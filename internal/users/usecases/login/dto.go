package login

type DTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type Response struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
