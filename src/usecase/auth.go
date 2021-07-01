package usecase

import (
	"errors"

	"github.com/creamyshit/gologin/domain"
	"github.com/creamyshit/gologin/src/model"
	"github.com/creamyshit/gologin/utils"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	authRepo domain.AuthRepository
}

func NewAuthUsecase(a domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{
		authRepo: a,
	}
}

func (a *AuthUsecase) ForgotPassword(forgot *model.ForgotPasswordPayload) (bool, error) {

	loginData := &model.Auth{
		Username: forgot.Username,
	}

	//find signin credential in user table db
	_, err := a.authRepo.GetUserbyUsername(loginData)

	//if failed finding user
	if err != nil {
		return false, err
	}
	return false, nil
}

func (a *AuthUsecase) Signin(signin *model.SignUpPayload) (*model.Auth, error) {

	//store signin credential <username> into user model struct
	loginData := &model.Auth{
		Username: signin.Username,
	}

	//find signin credential in user table db
	res, err := a.authRepo.GetUserbyUsername(loginData)

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

func (a *AuthUsecase) Signup(signup *model.SignUpPayload) (*model.Auth, error) {

	uuID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	salt := utils.GenerateRandomSalt(16)

	udata := &model.Auth{
		Id:       uuID,
		Username: signup.Username,
		Password: utils.HashPassword(signup.Password, salt),
		Salt:     salt,
	}

	res, err := a.authRepo.Signup(udata)

	return utils.HideCredential(res), nil
}
