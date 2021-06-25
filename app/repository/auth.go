package repository

import (
	"github.com/creamyshit/gologin/domain"
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

func (m *AuthRepository) Signup(a *domain.Auth) (*domain.Auth, error) {

	if err := m.conn.Create(&a); err != nil {
		return nil, err.Error
	}

	return a, nil
}

func (m *AuthRepository) Signin(a *domain.Auth) (*domain.Auth, error) {

	if err := m.conn.Where("username = ?", a.Username).First(&a); err != nil {
		return nil, err.Error
	}
	return a, nil
}
