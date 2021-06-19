package repoauth

import (
	"github.com/creamyshit/gologin/domain/auth"
	"github.com/creamyshit/gologin/model"

	"gorm.io/gorm"
)

type Repository struct {
	conn *gorm.DB
}

func AuthRepository(conn *gorm.DB) auth.Repository {
	return &Repository{
		conn: conn,
	}
}

func (m *Repository) Signup(a *model.User) (*model.User, error) {
	err := m.conn.Create(&a)
	if err.Error != nil {
		return nil, err.Error
	}

	return a, nil
}

func (m *Repository) Signin(a *model.User) (*model.User, error) {

	err := m.conn.Where("username = ?", a.Username).First(&a)

	if err.Error != nil {
		return nil, err.Error
	}
	return a, nil
}
