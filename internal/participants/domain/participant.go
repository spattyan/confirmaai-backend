package domain

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	EventID uuid.UUID `json:"event_id" gorm:"type:uuid;not null"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	RoleID  uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
