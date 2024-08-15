package handlerutils

import (
	"encoding/json"
	"errors"
	"github.com/matevskial/chirpyx/common"
	"log"
	"net/http"
	"path"
	"strings"
)

type errorResponse struct {
	Error string `json:"error"`
}

func encodeResponse(payload interface{}) ([]byte, error) {
	response, err := json.Marshal(payload)
	return response, err
}

func respond(w http.ResponseWriter, contentType string, statusCode int, content string) {
	respondBytes(w, contentType, statusCode, []byte(content))
}

func respondBytes(w http.ResponseWriter, contentType string, statusCode int, content []byte) {
	w.WriteHeader(statusCode)
	if len(contentType) != 0 {
		w.Header().Set("Content-Type", contentType)
		_, err := w.Write(content)
		if err != nil {
			http.Error(w, "Internal server error", 500)
		}
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
	respondBytes(w, "application/json", statusCode, response)
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

func RespondWithInternalServerError(w http.ResponseWriter) {
	RespondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func RespondWithUnauthorized(w http.ResponseWriter) {
	RespondWithError(w, http.StatusUnauthorized, "unauthorized")
}

func RespondWithStatusCode(w http.ResponseWriter, statusCode int) {
	respondBytes(w, "", statusCode, []byte{})
}

func PostRequestPath(pathElements ...string) string {
	return requestPath(http.MethodPost, pathElements...)
}

func requestPath(method string, pathElements ...string) string {
	return strings.Join([]string{method, path.Join(pathElements...)}, " ")
}

func PutRequestPath(pathElements ...string) string {
	return requestPath(http.MethodPut, pathElements...)
}

func SetSorting(req *http.Request, sorting *common.Sorting) error {
	sortingDirection := req.URL.Query().Get("sort")
	if sortingDirection == "" {
		sorting.Direction = common.Asc
		return nil
	}
	if sortingDirection != common.Asc && sortingDirection != common.Desc {
		return errors.New("error parsing sorting query parameters")
	}
	sorting.Direction = common.SortingDirection(sortingDirection)
	return nil
}
