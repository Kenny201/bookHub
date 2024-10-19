package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorResponse представляет формат ответа об ошибке.
type ErrorResponse struct {
	Code   int         `json:"code"`
	Error  string      `json:"error"`
	Detail interface{} `json:"detail,omitempty"`
}

// respondWithError отправляет ответ об ошибке в формате JSON.
func respondWithError(w http.ResponseWriter, code int, errorMessage string, detail interface{}) {
	respondWithJSON(w, code, ErrorResponse{Code: code, Error: errorMessage, Detail: detail})
}

// respondWithJSON отправляет успешный ответ в формате JSON.
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if payload == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		// Здесь не нужно вызывать respondWithError снова, так как это приведет к повторному вызову WriteHeader.
		http.Error(w, fmt.Sprintf(`{"error": "failed to marshal response", "details": "%s"}`, err.Error()), http.StatusInternalServerError)
	}
}
