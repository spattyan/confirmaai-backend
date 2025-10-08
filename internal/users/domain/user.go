package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/internal/participants/domain"
	"gorm.io/gorm"
)

type User struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name  string    `json:"name" gorm:"not null"`
	Email string    `json:"email" gorm:"unique;not null"`
	Phone string    `json:"phone" gorm:"not null"`

	Participants []domain.Participant `json:"participants" gorm:"foreignKey:UserID"`

	Password string `json:"password" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	user.ID = id
	return
}
