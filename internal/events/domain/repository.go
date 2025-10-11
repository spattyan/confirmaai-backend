package domain

import "github.com/spattyan/confirmaai-backend/internal/participants/domain"

type Repository interface {
	Create(event *Event) error
	FindByID(id string) (*Event, error)
	Update(event *Event) error
	Delete(id string) error
	List() ([]Event, error)

	ListParticipantsByEventID(eventID string) ([]domain.Participant, error)

	CreateEventRole(role *EventRole) error
	FindEventRoleByID(id string) (*EventRole, error)
	UpdateEventRole(role *EventRole) error
	DeleteEventRole(id string) error
	ListEventRolesByEventID(eventID string) ([]EventRole, error)
}
