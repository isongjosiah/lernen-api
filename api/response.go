package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//TODO(josiah): restructure this file later
type ErrorResponse struct {
	ErrorMessage string
}

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	message := ErrorResponse{ErrorMessage: err.Error()}
	if err != nil {
		JSONResponse(w, statusCode, message)
		return
	}
	JSONResponse(w, http.StatusBadRequest, nil)
}
