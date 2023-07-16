package main

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/configs"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/infrastructure/database"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/infrastructure/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
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

	router.Route("/products", func(router chi.Router) {
		router.Use(jwtauth.Verifier(configs.TokenAuth))
		router.Use(jwtauth.Authenticator)
		router.Post("/", productHandler.CreateProduct)
		router.Get("/{id}", productHandler.GetProduct)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.DeleteProduct)
		router.Get("/", productHandler.GetProducts)
	})

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", router)

}
