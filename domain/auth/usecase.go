package auth

import "github.com/creamyshit/gologin/model"

type SignUpPayload struct {
	Username string
	Password string
}

type Usecase interface {
	Signup(*SignUpPayload) (*model.User, error)
	Signin(*SignUpPayload) (*model.User, error)
}
