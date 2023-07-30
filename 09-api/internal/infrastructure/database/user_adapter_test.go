package database

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
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

func TestUserAdapter_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.User{})
	user, _ := model.NewUser("Daniel Silva", "daniel@test.email2", "12345")
	adapter := NewUserAdapter(db)
	adapter.Save(user)

	userFound, err := adapter.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, userFound.Email, user.Email)
}
