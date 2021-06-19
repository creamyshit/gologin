package auth

import "github.com/creamyshit/gologin/model"

type Repository interface {
	Signup(a *model.User) (*model.User, error)
	Signin(a *model.User) (*model.User, error)
}
