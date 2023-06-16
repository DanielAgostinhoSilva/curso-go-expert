package database

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/api/82-criando-user-db/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestUserAdapter_Save(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.User{})

	user, err := model.NewUser("Daniel Silva", "daniel@test.email", "12345")
	if err != nil {
		t.Error(err)
	}

	adapter := NewUserAdapter(db)
	err = adapter.Save(user)
	assert.Nil(t, err)

	var userFound model.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
}
