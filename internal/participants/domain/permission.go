package domain

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Name string `json:"name"`
	Key  string `json:"key"`

	Active bool `json:"active"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
