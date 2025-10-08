package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/internal/participants/domain"
	"gorm.io/gorm"
)

type Event struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title            string    `json:"title"`
	Description      string    `json:"description" gorm:""`
	Location         string    `json:"location" gorm:"not null"`
	DateAndTime      time.Time `json:"date_and_time" gorm:"not null"`
	ParticipantLimit int       `json:"participant_limit" gorm:"not null"`

	Participants []domain.Participant `json:"participants" gorm:"foreignKey:EventID"`
	Roles        []EventRole          `json:"roles" gorm:"foreignKey:EventID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (event *Event) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	event.ID = id
	return
}
