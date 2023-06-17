package repository

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
)

type UserRepository interface {
	Save(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
