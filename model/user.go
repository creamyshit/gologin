package model

import (
	"github.com/google/uuid"
)

type User struct {
	Userid   uuid.UUID `gorm:"type:uuid;primary_key"`
	Username string    `gorm:"type:character(24);not null"`
	Password string    `gorm:"type:varchar(100);not null"`
	Salt     []byte    `gorm:"type:bytea;not null"`
}
