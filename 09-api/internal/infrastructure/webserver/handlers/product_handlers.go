package handlers

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/dto"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/infrastructure/database"
	"net/http"
)

type ProductHandler struct {
	ProductDB database.ProductAdapter
}

func NewProductHandler(db database.ProductAdapter) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productInput dto.ProductInput
	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := model.NewProduct(productInput.Name, productInput.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Save(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
