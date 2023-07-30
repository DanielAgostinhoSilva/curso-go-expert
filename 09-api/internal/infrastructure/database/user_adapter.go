package database

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"gorm.io/gorm"
)

type UserAdapter struct {
	DB *gorm.DB
}

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{DB: db}
}

func (adapter *UserAdapter) Save(user *model.User) error {
	return adapter.DB.Create(user).Error
}

func (adapter *UserAdapter) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := adapter.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
