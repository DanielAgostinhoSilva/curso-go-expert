package main

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/configs"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/infrastructure/database"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/infrastructure/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	configs, err := configs.LoadConfig("./cmd/server/.env")
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

	userAdapter := database.NewUserAdapter(db)
	userHandler := handlers.NewUserHandler(*userAdapter, configs.TokenAuth, configs.JWTExpiresIn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products/{id}", productHandler.GetProduct)
	router.Put("/products/{id}", productHandler.UpdateProduct)
	router.Delete("/products/{id}", productHandler.DeleteProduct)
	router.Get("/products", productHandler.GetProducts)

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", router)

}
