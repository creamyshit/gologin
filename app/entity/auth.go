package entity

import (
	"errors"

	"github.com/creamyshit/gologin/domain"
	"github.com/creamyshit/gologin/utils"
	"github.com/google/uuid"
)

type AuthEntity struct {
	authRepo domain.AuthRepository
}

func NewAuthEntity(a domain.AuthRepository) domain.AuthEntity {
	return &AuthEntity{
		authRepo: a,
	}
}

func (a *AuthEntity) Signin(signin *domain.SignUpPayload) (*domain.Auth, error) {

	//store signin credential <username> into user model struct
	loginData := &domain.Auth{
		Username: signin.Username,
	}

	//find signin credential in user table db
	res, err := a.authRepo.Signin(loginData)

	//if failed finding user
	if err != nil {
		return nil, err
	}

	//if user found , check stored hashedpassword with signin credential password
	passMatch := utils.DoPasswordsMatch(res.Password, signin.Password, res.Salt)

	//if both password not same will return error
	if !passMatch {
		return nil, errors.New("auth failed")
	}

	//if both password match return fulldb result data
	//hidecredential used for reassign credential data into empty
	return utils.HideCredential(res), nil
}

func (a *AuthEntity) Signup(signup *domain.SignUpPayload) (*domain.Auth, error) {

	uuID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	salt := utils.GenerateRandomSalt(16)

	udata := &domain.Auth{
		Id:       uuID,
		Username: signup.Username,
		Password: utils.HashPassword(signup.Password, salt),
		Salt:     salt,
	}

	res, err := a.authRepo.Signup(udata)

	return utils.HideCredential(res), nil
}
