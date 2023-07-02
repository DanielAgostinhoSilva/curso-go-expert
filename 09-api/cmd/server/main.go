package main

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/infrastructure/database"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/infrastructure/webserver/handlers"
	"github.com/DanielAgostinhoSilva/curso-go-expert/api/77-criando-arquivos-de-configuracao/configs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig("./cmd/server/.env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Product{}, &model.User{})

	productAdapter := database.NewProductAdapter(db)
	productHandler := handlers.NewProductHandler(*productAdapter)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", router)

}
