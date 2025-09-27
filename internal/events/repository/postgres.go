package repository

import (
	"errors"
	"log"

	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"gorm.io/gorm"
)

type gormRepository struct {
	database *gorm.DB
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
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func NewGormRepository(database *gorm.DB) domain.Repository {
	return &gormRepository{
		database: database,
	}
}
