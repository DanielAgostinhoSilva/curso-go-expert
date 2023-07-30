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
	"strconv"
)

type ProductHandler struct {
	ProductDB database.ProductAdapter
}

func NewProductHandler(db database.ProductAdapter) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct godoc
// @Summary      Create Product
// @Description  Create Product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request body dto.ProductInput true "user request"
// @Success      201
// @Failure      400  {object} Error
// @Failure      500  {object} Error
// @Router       /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productInput dto.ProductInput
	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		HandlerError(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := model.NewProduct(productInput.Name, productInput.Price)
	if err != nil {
		HandlerError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Save(product)
	if err != nil {
		HandlerError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary      Get a Product
// @Description  Get a Product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 id	path string	true "product ID" Format(uuid)
// @Success      200 {object} dto.ProductModel
// @Failure      401 {object} Error
// @Failure      404 {object} Error
// @Failure      500 {object} Error
// @Router       /products/{id} [get]
// @Security ApiKeyAuth
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

// GetProducts godoc
// @Summary      List all Product
// @Description  Get all Product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 page	query	string	false	"page number"
// @Param		 limit	query	string	false	"limit"
// @Success      200 {array} dto.ProductModel
// @Failure      401 {object} Error
// @Failure      404 {object} Error
// @Failure      500 {object} Error
// @Router       /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	sort := r.URL.Query().Get("sort")
	productsFound, err := h.ProductDB.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var productsModel []dto.ProductModel
	err = copier.Copy(&productsModel, &productsFound)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productsModel)
}

// UpdateProduct godoc
// @Summary      Update a Product
// @Description  Update a Product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path string true "product ID" Format(uuid)
// @Param        request body dto.ProductInput true "user request"
// @Success      200
// @Failure      400  {object} Error
// @Failure      401  {object} Error
// @Failure      404  {object} Error
// @Failure      500  {object} Error
// @Router       /products/{id} [put]
// @Security ApiKeyAuth
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

// DeleteProduct godoc
// @Summary      Delete a Product
// @Description  Delete a Product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path string true "product ID" Format(uuid)
// @Success      204
// @Failure      400  {object} Error
// @Failure      401  {object} Error
// @Failure      404  {object} Error
// @Failure      500  {object} Error
// @Router       /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
