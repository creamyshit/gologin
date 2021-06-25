package domain

import (
	"github.com/google/uuid"
)

type Auth struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key"`
	Username string    `gorm:"type:character(24);not null"`
	Password string    `gorm:"type:varchar(100);not null"`
	Salt     []byte    `gorm:"type:bytea;not null"`
}

type AuthRepository interface {
	Signup(a *Auth) (*Auth, error)
	Signin(a *Auth) (*Auth, error)
}

type SignUpPayload struct {
	Username string `validate:"required,min=5,max=25"`
	Password string `validate:"required,min=8,max=25"`
}

type AuthEntity interface {
	Signup(*SignUpPayload) (*Auth, error)
	Signin(*SignUpPayload) (*Auth, error)
}
