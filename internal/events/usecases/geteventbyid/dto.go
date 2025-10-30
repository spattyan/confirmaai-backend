package geteventbyid

type DTO struct {
	Id string `json:"id" validate:"required,uuid"`
}

type Response struct {
	Event EventResponse `json:"event"`
}

type EventResponse struct {
	Title            string               `json:"title"`
	Description      string               `json:"description"`
	Location         string               `json:"location"`
	DateAndTime      string               `json:"date_and_time"`
	ParticipantLimit int                  `json:"participant_limit"`
	CreatedBy        string               `json:"created_by"`
	Participants     ParticipantsResponse `json:"participants"`
}

type ParticipantsResponse struct {
	Count int                        `json:"count"`
	List  []ParticipantsListResponse `json:"list_events"`
}

type ParticipantsListResponse struct {
	ID string `json:"id"`
}
