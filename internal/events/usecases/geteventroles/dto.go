package geteventroles

type DTO struct {
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
