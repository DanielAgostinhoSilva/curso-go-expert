package repository

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/api/82-criando-user-db/internal/domain/model"
)

type UserRepository interface {
	Save(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
