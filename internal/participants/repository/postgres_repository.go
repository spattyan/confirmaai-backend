package repository

import (
	"errors"

	eventDomain "github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/participants/domain"
	userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormRepository struct {
	database *gorm.DB
}

func (g gormRepository) CreatePermission(name, key string) error {
	permission := domain.Permission{
		Name:   name,
		Key:    key,
		Active: true,
	}

	err := g.database.Create(&permission)

	if err != nil {
		return errors.New("unable to createevent permission")
	}

	return nil
}

func (g gormRepository) GetPermissionByKey(key string) (*domain.Permission, error) {
	var permission domain.Permission

	err := g.database.First(&permission, "key = ?", key).Error

	if err != nil {
		return nil, errors.New("permission not found")
	}

	return &permission, nil
}

func (g gormRepository) Create(participant *domain.Participant) error {
	var event eventDomain.Event
	var user userDomain.User

	err := g.database.First(&event, "id = ?", participant.EventID).Error

	if err != nil {
		return errors.New("event not found")
	}

	err = g.database.First(&user, "id = ?", participant.UserID).Error

	if err != nil {
		return errors.New("user not found")
	}

	err = g.database.Create(&participant).Error

	if err != nil {
		return errors.New("failed to createevent participant")
	}
	return nil
}

func (g gormRepository) FindByID(id string) (*domain.Participant, error) {
	var participant domain.Participant
	err := g.database.Preload("User").Preload("Event").Preload("Permission").First(&participant, "id = ?", id).Error

	if err != nil {
		return nil, errors.New("participant not found")
	}

	return &participant, nil
}

func (g gormRepository) FindByUser(userID string) ([]domain.Participant, error) {
	var participants []domain.Participant
	err := g.database.Preload(clause.Associations).Where("user_id = ?", userID).Find(&participants).Error
	if err != nil {
		return nil, errors.New("failed to fetch participants")
	}
	return participants, nil
}

func (g gormRepository) FindByEvent(eventID string) ([]domain.Participant, error) {
	var participants []domain.Participant
	err := g.database.Preload(clause.Associations).Where("event_id = ?", eventID).Find(&participants).Error
	if err != nil {
		return nil, errors.New("failed to fetch participants")
	}
	return participants, nil
}

func (g gormRepository) FindByEventAndUser(eventID, userID string) (*domain.Participant, error) {
	var participant domain.Participant
	err := g.database.Preload(clause.Associations).Where("event_id = ? AND user_id = ?", eventID, userID).First(&participant).Error
	if err != nil {
		return nil, errors.New("participant not found")
	}
	return &participant, nil
}

func (g gormRepository) Update(participant *domain.Participant) error {
	err := g.database.Save(participant).Error

	if err != nil {
		return errors.New("failed to update participant")
	}

	return nil
}

func (g gormRepository) Delete(id string) error {
	err := g.database.Delete(&domain.Participant{}, "id = ?", id).Error

	if err != nil {
		return errors.New("failed to delete participant")
	}
	return nil
}

func NewGormRepository(database *gorm.DB) domain.Repository {
	return &gormRepository{
		database: database,
	}
}
