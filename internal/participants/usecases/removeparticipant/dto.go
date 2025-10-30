package removeparticipant

type DTO struct {
	ParticipantID string `json:"participant_id" validate:"required,uuid"`
}

type Response struct {
	ID string `json:"id"`
}
