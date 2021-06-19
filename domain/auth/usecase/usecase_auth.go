package usecase

import (
	"errors"

	"github.com/creamyshit/gologin/domain/auth"
	"github.com/creamyshit/gologin/helper"
	"github.com/creamyshit/gologin/model"

	"github.com/google/uuid"
)

type Usecase struct {
	authRepo auth.Repository
}

func AuthUsecase(a auth.Repository) auth.Usecase {
	return &Usecase{
		authRepo: a,
	}
}

func (a *Usecase) Signin(signin *auth.SignUpPayload) (*model.User, error) {

	//store signin credential <username> into user model struct
	loginData := &model.User{
		Username: signin.Username,
	}

	//find signin credential in user table db
	res, err := a.authRepo.Signin(loginData)

	//if failed finding user
	if err != nil {
		return nil, err
	}

	//if user found , check stored hashedpassword with signin credential password
	passMatch := helper.DoPasswordsMatch(res.Password, signin.Password, res.Salt)

	//if both password not same will return error
	if !passMatch {
		return nil, errors.New("auth failed")
	}

	//if both password match return fulldb result data
	//hidecredential used for reassign credential data into empty
	return helper.HideCredential(res), nil
}

func (a *Usecase) Signup(signup *auth.SignUpPayload) (*model.User, error) {

	uuID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	salt := helper.GenerateRandomSalt(16)

	udata := &model.User{
		Userid:   uuID,
		Username: signup.Username,
		Password: helper.HashPassword(signup.Password, salt),
		Salt:     salt,
	}

	res, err := a.authRepo.Signup(udata)

	return res, nil
}
