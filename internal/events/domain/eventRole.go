package domain

import (
	"time"

	"github.com/google/uuid"
)

type EventRole struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	EventID uuid.UUID `json:"event_id" gorm:"type:uuid;not null"`
	Event   Event     `json:"event" gorm:"foreignKey:EventID"`

	Name  string `json:"name" gorm:"not null"`
	Slots int    `json:"slots" gorm:"not null"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
