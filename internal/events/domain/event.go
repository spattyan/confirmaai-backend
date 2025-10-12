package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/internal/participants/domain"
	userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type Event struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title            string    `json:"title"`
	Description      string    `json:"description" gorm:""`
	Location         string    `json:"location" gorm:"not null"`
	DateAndTime      time.Time `json:"date_and_time" gorm:"not null"`
	ParticipantLimit int       `json:"participant_limit" gorm:"not null"`

	CreatedByID uuid.UUID       `json:"created_by_id" gorm:"type:uuid;not null"`
	CreatedBy   userDomain.User `json:"created_by,omitempty" gorm:"foreignKey:CreatedByID"`

	Participants []domain.Participant `json:"participants" gorm:"foreignKey:EventID"`
	Roles        []EventRole          `json:"roles" gorm:"foreignKey:EventID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
