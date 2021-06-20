package auth

import "github.com/creamyshit/gologin/model"

type SignUpPayload struct {
	Username string `validate:"required,min=5,max=25"`
	Password string `validate:"required,min=8,max=25"`
}

type Usecase interface {
	Signup(*SignUpPayload) (*model.User, error)
	Signin(*SignUpPayload) (*model.User, error)
}
