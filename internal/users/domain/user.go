package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/internal/participants/domain"
)

type User struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name  string    `json:"name,omitempty" gorm:"not null"`
	Email string    `json:"email,omitempty" gorm:"unique;not null"`
	Phone string    `json:"phone,omitempty" gorm:"not null"`

	Participants []domain.Participant `json:"participants,omitempty" gorm:"foreignKey:UserID"`

	Password string `json:"password,omitempty" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
