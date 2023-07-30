package dto

type JWTUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
