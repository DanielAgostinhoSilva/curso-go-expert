package handlers

import (
	"encoding/json"
	"net/http"
)

func HandlerError(w http.ResponseWriter, msg string, httpStatus int) {
	w.WriteHeader(httpStatus)
	error := Error{Message: msg}
	json.NewEncoder(w).Encode(error)
}
