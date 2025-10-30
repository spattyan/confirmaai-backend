package listevents

type Response struct {
	Count  int              `json:"count"`
	Events []ResponseObject `json:"events"`
}
type ResponseObject struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	Location         string `json:"location"`
	DateAndTime      string `json:"date_and_time"`
	ParticipantLimit int    `json:"participant_limit"`
}

type DTO struct {
}
