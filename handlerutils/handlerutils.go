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

func RespondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	errorResponseDto := errorResponse{Error: errorMessage}
	RespondWithJson(w, statusCode, errorResponseDto)
}
