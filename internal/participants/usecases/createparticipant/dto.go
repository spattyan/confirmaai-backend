package createparticipant

type DTO struct {
	EventID string `json:"event_id" validate:"required,uuid"`
	UserID  string `json:"user_id" validate:"required,uuid"`
	RoleID  string `json:"role_id" validate:"required,uuid"`
}

type Response struct {
	ID string `json:"id"`
}
