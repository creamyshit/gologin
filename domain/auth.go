package domain

import "github.com/creamyshit/gologin/src/model"

type AuthUsecase interface {
	Signup(*model.SignUpPayload) (*model.Auth, error)
	Signin(*model.SignUpPayload) (*model.Auth, error)
	ForgotPassword(*model.ForgotPasswordPayload) (bool, error)
}

type AuthRepository interface {
	Signup(a *model.Auth) (*model.Auth, error)
	GetUserbyUsername(a *model.Auth) (*model.Auth, error)
}
