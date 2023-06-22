package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `gorm:"index"`
	Id        uuid.UUID  `gorm:"primaryKey"`
	Role      string
	Password  string
}
