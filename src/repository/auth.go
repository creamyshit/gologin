package repository

import (
	"github.com/creamyshit/gologin/domain"
	"github.com/creamyshit/gologin/src/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	conn *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) domain.AuthRepository {
	return &AuthRepository{
		conn: conn,
	}
}

func (m *AuthRepository) Signup(a *model.Auth) (*model.Auth, error) {

	if err := m.conn.Create(&a); err != nil {
		return nil, err.Error
	}

	return a, nil
}

func (m *AuthRepository) GetUserbyUsername(a *model.Auth) (*model.Auth, error) {

	if err := m.conn.Where("username = ?", a.Username).First(&a); err != nil {
		return nil, err.Error
	}
	return a, nil
}
