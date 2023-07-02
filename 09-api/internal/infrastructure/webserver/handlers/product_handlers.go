package handlers

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/dto"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var productModel dto.ProductModel
	err = copier.Copy(&productModel, &product)

	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var productInput dto.ProductInput
	err = json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	productFound, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = copier.Copy(&productFound, &productInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.ProductDB.Update(productFound)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
