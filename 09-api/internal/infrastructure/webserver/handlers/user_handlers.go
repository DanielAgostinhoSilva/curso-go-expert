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

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var jwtUserInput dto.JWTUserInput
	err := json.NewDecoder(r.Body).Decode(&jwtUserInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.UserDB.FindByEmail(jwtUserInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !user.ValidatePassword(jwtUserInput.Password) {
		w.WriteHeader(http.StatusUnauthorized)
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
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	user, err := model.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.UserDB.Save(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
