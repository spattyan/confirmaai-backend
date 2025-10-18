package domain

import (
	"time"

	"github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type EventInvite struct {
	ID      string      `json:"id"`
	EventID string      `json:"event_id"`
	Event   domain.User `json:"event,omitempty" gorm:"foreignKey:EventID"`

	Active bool `json:"active"`

	InviteHash string `json:"invite_hash"`

	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
