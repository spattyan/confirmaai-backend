package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title            string    `json:"title"`
	Description      string    `json:"description" gorm:""`
	Location         string    `gorm:"not null"`
	DateAndTime      time.Time `gorm:"not null"`
	ParticipantLimit int       `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *Event) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	user.ID = id
	return
}
