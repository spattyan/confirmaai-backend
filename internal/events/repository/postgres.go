package repository

import (
	"errors"
	"log"

	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	participantDomain "github.com/spattyan/confirmaai-backend/internal/participants/domain"
	"gorm.io/gorm"
)

type gormRepository struct {
	database *gorm.DB
}

func (g gormRepository) ListParticipantsByEventID(eventID string) ([]participantDomain.Participant, error) {
	var event domain.Event
	result := g.database.Preload("Participants").First(&event, "id = ?", eventID)

	if result.Error != nil {
		log.Printf("Error fetching participants: %v", result.Error)
		return nil, errors.New("failed to fetch participants")
	}

	return event.Participants, nil
}

func (g gormRepository) CreateEventRole(role *domain.EventRole) error {
	err := g.database.Create(&role).Error

	if err != nil {
		log.Printf("Error creating event role: %v", err)
		return errors.New("failed to create event role")
	}

	return nil
}

func (g gormRepository) FindEventRoleByID(id string) (*domain.EventRole, error) {
	var role domain.EventRole
	result := g.database.First(&role, "id = ?", id)

	if result.Error != nil {
		log.Printf("Error fetching event role: %v", result.Error)
		return nil, errors.New("failed to fetch event role")
	}
	return &role, nil
}

func (g gormRepository) UpdateEventRole(role *domain.EventRole) error {
	err := g.database.Save(&role).Error
	if err != nil {
		return errors.New("failed to update event role")
	}
	return nil
}

func (g gormRepository) DeleteEventRole(id string) error {
	err := g.database.Delete(&domain.EventRole{}, "id = ?", id).Error
	if err != nil {
		return errors.New("failed to delete event role")
	}
	return nil
}

func (g gormRepository) ListEventRolesByEventID(eventID string) ([]domain.EventRole, error) {
	var roles []domain.EventRole
	result := g.database.Where("event_id = ?", eventID).Find(&roles)

	if result.Error != nil {
		log.Printf("Error fetching event roles: %v", result.Error)
		return nil, errors.New("failed to fetch event roles")
	}

	return roles, nil
}

func (g gormRepository) Create(event *domain.Event) error {
	err := g.database.Create(&event).Error

	if err != nil {
		log.Printf("Error creating event: %v", err)
		return errors.New("failed to create event")
	}

	return nil
}

func (g gormRepository) FindByID(id string) (*domain.Event, error) {
	var event domain.Event
	result := g.database.
		Preload("Participants", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "event_id", "user_id", "role_id")
		}).
		First(&event, "id = ?", id)

	if result.Error != nil {
		log.Printf("Error fetching events: %v", result.Error)
		return nil, errors.New("failed to fetch events")
	}

	return &event, nil
}

func (g gormRepository) Update(event *domain.Event) error {
	//TODO implement me
	panic("implement me")
}

func (g gormRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (g gormRepository) List() ([]domain.Event, error) {
	var events []domain.Event
	result := g.database.Find(&events)

	if result.Error != nil {
		log.Printf("Error fetching events: %v", result.Error)
		return nil, errors.New("failed to fetch events")
	}

	return events, nil
}

func NewGormRepository(database *gorm.DB) domain.Repository {
	return &gormRepository{
		database: database,
	}
}
