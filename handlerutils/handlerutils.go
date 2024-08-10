package handlerutils

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func encodeResponse(payload interface{}) ([]byte, error) {
	response, err := json.Marshal(payload)
	return response, err
}

func respond(w http.ResponseWriter, contentType string, statusCode int, content string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(content))
	if err != nil {
		http.Error(w, "Internal server error", 500)
	}
}

func RespondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := encodeResponse(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RespondWithText(w http.ResponseWriter, statusCode int, text string) {
	respond(w, "text/plain; charset=utf-8", statusCode, text)
}

func RespondWithHtml(w http.ResponseWriter, statusCode int, html string) {
	respond(w, "text/html", statusCode, html)
}

func RespondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	errorResponseDto := errorResponse{Error: errorMessage}
	RespondWithJson(w, statusCode, errorResponseDto)
}
