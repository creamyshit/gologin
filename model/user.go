package model

import "github.com/google/uuid"

type User struct {
	Userid   uuid.UUID
	Username string
	Password string
	Salt     []byte
}
