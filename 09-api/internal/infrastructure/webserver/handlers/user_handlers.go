package handlers

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/dto"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/infrastructure/database"
	"github.com/go-chi/jwtauth"
	"net/http"
	"time"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserAdapter
}

func NewUserHandler(db database.UserAdapter) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body dto.JWTUserInput true "user credentials"
// @Success      200 {object} dto.TokenModel
// @Failure      401 {object} Error
// @Failure      404 {object} Error
// @Failure      500 {object} Error
// @Router       /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var jwtUserInput dto.JWTUserInput
	err := json.NewDecoder(r.Body).Decode(&jwtUserInput)
	if err != nil {
		handlerError(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.UserDB.FindByEmail(jwtUserInput.Email)
	if err != nil {
		handlerError(w, err.Error(), http.StatusNotFound)
		return
	}
	if !user.ValidatePassword(jwtUserInput.Password) {
		handlerError(w, "Invalid user or password", http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	token := dto.TokenModel{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

// CreateUser Create user godoc
// @Summary      Create User
// @Description  Create User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body dto.UserInput true "user request"
// @Success      201
// @Failure      500  {object} Error
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		handlerError(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		handlerError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UserDB.Save(user)
	if err != nil {
		handlerError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func handlerError(w http.ResponseWriter, msg string, httpStatus int) {
	w.WriteHeader(httpStatus)
	error := Error{Message: msg}
	json.NewEncoder(w).Encode(error)
}
